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

func GetFriendListForStrengthRequesst(data []byte,account *kaccount.Account,conn net.Conn) error {

	ms := &kpb.GetFriendListForStrengthRequesst{}
	err := proto.Unmarshal(data, ms)
	if err!= nil{
		klog.Klog.Println(err)
		return err
	}

	lists := []uint64{101,102,103}
	GetFriendListForStrengthResponse(kother.RESCODE_Success,lists,conn)
	return nil
}
func GetFriendListForStrengthResponse(code int,lists []uint64,conn net.Conn)  {
	mResp := &kpb.GetFriendListForStrengthResponse{}
	mResp.ErrorCode = proto.Int32(int32(code))
	mResp.AccountIds =lists
	data,err := proto.Marshal(mResp)
	if err != nil{
		klog.Klog.Println(err)
		return
	}

	sendBuf := knet.SetPackage(kpb.MSGTYPE_GetFriendListForStrengthResponse,data)

	if _ , err := conn.Write(sendBuf); err != nil {
		klog.Klog.Println(err)
		return
	}
	return
}


