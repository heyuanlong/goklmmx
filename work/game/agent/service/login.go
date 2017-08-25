package service

import (
	"net"
	"github.com/golang/protobuf/proto"

	knet "goklmmx/lib/net"
	klog "goklmmx/lib/log"
	kpb "goklmmx/lib/pb"
	kredis "goklmmx/lib/db/redis"
	kother "goklmmx/lib/other"

	kaccount "goklmmx/work/game/common/account"

)

func GameServerLoginRequest(data []byte,account *kaccount.Account,conn net.Conn) (error,bool) {

	ms := &kpb.GameServerLoginRequest{}
	err := proto.Unmarshal(data, ms)
	if err!= nil{
		klog.Klog.Println(err)
		return err,false
	}
	code , err := kredis.GetAuth(  int(ms.GetAccountId()) )
	if err!= nil {
		klog.Klog.Println("GetAuth fail")
		GameServerLoginResponse(kother.RESCODE_LOGIN_AUTH_FAIL,conn)
		return nil,false
	}
	if code != ms.GetAuthCode(){
		klog.Klog.Println("auth code is fail")
		GameServerLoginResponse(kother.RESCODE_LOGIN_AUTH_FAIL,conn)
		return nil,false
	}
	account.AccountId = int(ms.GetAccountId())
	GameServerLoginResponse(kother.RESCODE_Success,conn)
	return nil,true
}
func GameServerLoginResponse(code int,conn net.Conn)  {
	mResp := &kpb.GameServerLoginResponse{}
	mResp.ErrorCode = proto.Int32(int32(code))
	data,err := proto.Marshal(mResp)
	if err != nil{
		klog.Klog.Println(err)
		return
	}
	sendBuf := knet.SetPackage(kpb.MSGTYPE_GameServerLoginResponse,data)

	if _ , err := conn.Write(sendBuf); err != nil {
		klog.Klog.Println(err)
		return
	}
	return
}
