package main

import (
	"net"
	"errors"

	kpb "goklmmx/lib/pb"
	kaccount "goklmmx/work/game/common/account"
	kservice "goklmmx/work/game/agent/service"
)

type Agent struct {
	conn net.Conn
	accout *kaccount.Account
	isLogin bool
}

func NewAgent(conn net.Conn) *Agent {
	agent := new(Agent)
	agent.conn = conn
	agent.accout = kaccount.NewAccount()
	agent.isLogin = false

	return agent
}

func (agent *Agent) DealPackage(msgType int ,pBuf []byte) error {
	if kpb.MSGTYPE_GameServerLoginRequest == msgType {
		err,isLogin := kservice.GameServerLoginRequest(pBuf,agent.accout,agent.conn)
		if err == nil {
			agent.isLogin = isLogin
		}
		return err
	}
	if agent.IsLogin() == false {
		return errors.New("Haven't login")
	}

	if kpb.MSGTYPE_SyncBaseInfoRequest == msgType {
		return kservice.SyncBaseInfoRequest(pBuf,agent.accout,agent.conn)
	}
	if kpb.MSGTYPE_PingRequest == msgType {
		return kservice.PingRequest(pBuf,agent.accout,agent.conn)
	}
	if kpb.MSGTYPE_GetFriendListForStrengthRequesst == msgType {
		return kservice.GetFriendListForStrengthRequesst(pBuf,agent.accout,agent.conn)
	}
	if kpb.MSGTYPE_SendGameResumeRequest == msgType {
		return kservice.SendGameResumeRequest(pBuf,agent.accout,agent.conn)
	}
	return nil
}
func (agent *Agent) IsLogin() bool {
	return agent.isLogin
}

