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

func SendGameResumeRequest(data []byte,account *kaccount.Account,conn net.Conn) error {

	ms := &kpb.SendGameResumeRequest{}
	err := proto.Unmarshal(data, ms)
	if err!= nil{
		klog.Klog.Println(err)
		return err
	}
	resume := ms.GetData()

	klog.Klog.Println("SendGameResumeRequest")
	klog.Klog.Println(resume.String())

	SendGameResumeResponse(kother.RESCODE_Success,conn)
	return nil
}
func SendGameResumeResponse(code int,conn net.Conn)  {
	mResp := &kpb.SendGameResumeResponse{}
	mResp.ErrorCode = proto.Int32(int32(code))
	data,err := proto.Marshal(mResp)
	if err != nil{
		klog.Klog.Println(err)
		return
	}
	sendBuf := knet.SetPackage(kpb.MSGTYPE_SendGameResumeResponse,data)

	if _ , err := conn.Write(sendBuf); err != nil {
		klog.Klog.Println(err)
		return
	}
	return
}


//-----------------------------------------------------------------------------------------
