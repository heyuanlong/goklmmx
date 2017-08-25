package service

import (
	"net"
	"github.com/golang/protobuf/proto"

	kaccount "goklmmx/work/game/common/account"
	kpb "goklmmx/lib/pb"
	klog "goklmmx/lib/log"
	kother "goklmmx/lib/other"
	knet "goklmmx/lib/net"

)

func SyncBaseInfoRequest(data []byte,account *kaccount.Account,conn net.Conn) error {

	ms := &kpb.SyncBaseInfoRequest{}
	err := proto.Unmarshal(data, ms)
	if err!= nil{
		klog.Klog.Println(err)
		return err
	}
	klog.Klog.Println(ms.GetNickname())
	klog.Klog.Println(ms.GetHeadImage())

	SyncBaseInfoResponse(kother.RESCODE_Success,conn)
	return nil
}
func SyncBaseInfoResponse(code int,conn net.Conn)  {
	mResp := &kpb.SyncBaseInfoResponse{}
	mResp.ErrorCode = proto.Int32(int32(code))
	data,err := proto.Marshal(mResp)
	if err != nil{
		klog.Klog.Println(err)
		return
	}
	sendBuf := knet.SetPackage(kpb.MSGTYPE_SyncBaseInfoResponse,data)

	if _ , err := conn.Write(sendBuf); err != nil {
		klog.Klog.Println(err)
		return
	}
	return
}


//-----------------------------------------------------------------------------------------
func PingRequest(data []byte,account *kaccount.Account,conn net.Conn) error {

	ms := &kpb.PingRequest{}
	err := proto.Unmarshal(data, ms)
	if err!= nil{
		klog.Klog.Println(err)
		return err
	}
	klog.Klog.Println(ms.GetClientTime())
	PingResponse(kother.RESCODE_Success,ms.GetClientTime(),conn)
	return nil
}
func PingResponse(code int,clientTime int32 ,conn net.Conn)  {
	mResp := &kpb.PingResponse{}
	mResp.ErrorCode = proto.Int32(int32(code))
	data,err := proto.Marshal(mResp)
	if err != nil{
		klog.Klog.Println(err)
		return
	}
	sendBuf := knet.SetPackage(kpb.MSGTYPE_PingResponse,data)

	if _ , err := conn.Write(sendBuf); err != nil {
		klog.Klog.Println(err)
		return
	}
	return
}
