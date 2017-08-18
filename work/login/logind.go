package main

import (
	"fmt"
	"net"
	"time"
	"errors"
	"github.com/golang/protobuf/proto"


	klog "goklmmx/lib/log"
	knet "goklmmx/lib/net"
	kpb "goklmmx/lib/pb"
	kmysql "goklmmx/lib/db/mysql"
	kredis "goklmmx/lib/db/redis"
	kutils "goklmmx/lib/utils"
	kother "goklmmx/lib/other"
)

const (
	G_MSG_SIZE_MAX    = 1024
	G_BUF_SIZE_MAX    = 65536
)

func HandleClient(conn net.Conn)  {
	klog.Klog.Println("HandleClient")
	defer conn.Close()
	var bufBuf = make([]byte,0)
	var msgBuf = make([]byte, G_MSG_SIZE_MAX)
	for  {
		conn.SetReadDeadline(time.Now().Add(kother.FD_TIMEOUT_SECOND * time.Second))
		n , err := conn.Read(msgBuf)
		if err!= nil{
			if nerr, ok := err.(*net.OpError); ok && nerr.Timeout() {
				klog.Klog.Println("timeout")
				return
			}else {
				klog.Klog.Println("read close or fail")
				return
			}
		}
		if (len(bufBuf) + n ) >  G_BUF_SIZE_MAX {
			klog.Klog.Println("buf too big")
			return
		}
		bufBuf = append(bufBuf,msgBuf[0:n]...)
		msgLen,msgType,pBuf  := knet.ParsePackage(bufBuf)
		if msgLen != 0 {
			sendBuf ,err := DealPackage(msgType,pBuf)
			if err != nil{
				klog.Klog.Println(err)
				return
			}else{
				klog.Klog.Println("send len:",len(sendBuf))
				conn.Write(sendBuf)
			}
			bufBuf = bufBuf[msgLen:]
		}
	}
}

func DealPackage(msgType int ,pBuf []byte)( []byte,error ){
	if kpb.MSGTYPE_LoginRequest == msgType {
		bufBytes , err := DealLogin(pBuf)
		if err != nil {
			return nil , err
		}
		return knet.SetPackage(kpb.MSGTYPE_LoginResponse,bufBytes),nil
	}
	return nil,errors.New(fmt.Sprintf("this msgid is invalid:%d",msgType))
}

func DealLogin(data []byte) ( []byte ,  error) {
	ms := &kpb.LoginRequest{}
	err := proto.Unmarshal(data, ms)
	if err!= nil{
		klog.Klog.Println(err)
		return loginFailMsg(kother.RESCODE_Fail)
	}
	accountId := kmysql.SelectAccount(ms.GetDeviceId())
	if accountId == 0 {
		accountId = kredis.GetNextAccountId()
		if accountId == 0 {
			klog.Klog.Println("GetNextAccountId fail")
			return loginFailMsg(kother.RESCODE_Fail)
		}
		if err := kmysql.CreateAccount(accountId,ms.GetDeviceId()) ; err != nil{
			klog.Klog.Println(err)
			return loginFailMsg(kother.RESCODE_Fail)
		}
	}

	authCode := kutils.Md5( fmt.Sprintf("%d_%s",accountId,kutils.GetRandomString(5)) )
	if err := kredis.SetAuth(accountId,authCode) ; err != nil {
		klog.Klog.Println(err)
		return loginFailMsg(kother.RESCODE_Fail)
	}

	mResp := &kpb.LoginResponse{}
	mResp.ErrorCode = proto.Int32(int32( kother.RESCODE_Success))
	mResp.AccountId =  proto.Int64(int64(accountId))
	mResp.AuthCode = proto.String(authCode)
	mResp.Ip =  proto.String("127.0.0.1")
	mResp.Port =  proto.Int32(9020)

	data,err_ := proto.Marshal(mResp)
	if err_ != nil{
		klog.Klog.Println(err_)
		return nil ,err_
	}
	klog.Klog.Println(mResp)
	return data ,nil
}

func loginFailMsg( errorCode int ) ([]byte ,error ) {
	mResp := &kpb.LoginResponse{}
	mResp.ErrorCode = proto.Int32(int32( kother.RESCODE_Fail))
	mResp.AccountId =  proto.Int64(int64(0))
	mResp.AuthCode = proto.String("")
	mResp.Ip =  proto.String("")
	mResp.Port =  proto.Int32(0)

	data,err_ := proto.Marshal(mResp)
	if err_ != nil{
		klog.Klog.Println(err_)
		return nil ,err_
	}
	return data ,nil
}

