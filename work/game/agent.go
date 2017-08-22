package main

import (
	"net"
	"github.com/golang/protobuf/proto"

	knet "goklmmx/lib/net"
	klog "goklmmx/lib/log"
	kpb "goklmmx/lib/pb"
	kredis "goklmmx/lib/db/redis"
	kother "goklmmx/lib/other"
)

type Agent struct {
	conn net.Conn
	accoutId int
	isLogin bool
}

func NewAgent(conn net.Conn) *Agent {
	agent := new(Agent)
	agent.conn = conn

	return agent
}

func (agent *Agent) DealPackage(msgType int ,pBuf []byte) error {
	if kpb.MSGTYPE_GameServerLoginRequest == msgType {
		return agent.DealLogin(pBuf)
	}
	return nil
}
func (agent *Agent) IsLogin() bool {
	return agent.isLogin
}


func (agent *Agent) DealLogin(data []byte) error {
	ms := &kpb.GameServerLoginRequest{}
	err := proto.Unmarshal(data, ms)
	if err!= nil{
		klog.Klog.Println(err)
		return err
	}
	code , err := kredis.GetAuth(  int(ms.GetAccountId()) )
	if err!= nil {
		klog.Klog.Println("GetAuth fail")
		agent.DealLoginResp(kother.RESCODE_LOGIN_AUTH_FAIL)
		return nil
	}
	if code != ms.GetAuthCode(){
		klog.Klog.Println("auth code is fail")
		agent.DealLoginResp(kother.RESCODE_LOGIN_AUTH_FAIL)
		return nil
	}
	agent.isLogin = true
	agent.accoutId = int(ms.GetAccountId())
	agent.DealLoginResp(kother.RESCODE_Success)
	return nil
}
func (agent *Agent) DealLoginResp(code int) {
	mResp := &kpb.GameServerLoginResponse{}
	mResp.ErrorCode = proto.Int32(int32(code))
	data,err := proto.Marshal(mResp)
	if err != nil{
		klog.Klog.Println(err)
		return
	}
	sendBuf := knet.SetPackage(kpb.MSGTYPE_GameServerLoginResponse,data)

	if _ , err := agent.conn.Write(sendBuf); err != nil {
		klog.Klog.Println(err)
		return
	}
	return
}