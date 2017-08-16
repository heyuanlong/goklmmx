package main

import (
	"net"
	"time"
	"github.com/golang/protobuf/proto"


	klog "goklmmx/lib/log"
	knet "goklmmx/lib/net"
	kpb "goklmmx/lib/pb"
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
		conn.SetReadDeadline(time.Now().Add(20 * time.Second))
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
		klog.Klog.Println(string(bufBuf))
		klog.Klog.Println(len(bufBuf))
		msgLen,msgType,pBuf  := knet.ParsePackage(bufBuf)
		if msgLen != 0 {
			if err := DealPackage(msgType,pBuf) ; err != nil{
				klog.Klog.Println(err)
				return
			}
			bufBuf = bufBuf[msgLen:]
		}
	}
}

func DealPackage(msgType int ,pBuf []byte) error {
	klog.Klog.Println(msgType,pBuf)
	if kpb.MSGTYPE["LoginRequest"] == msgType {
		DealLogin(pBuf)
	}
	return nil
}

func DealLogin(data []byte) error {
	ms := &kpb.LoginRequest{}
	err := proto.Unmarshal(data, ms)
	if err!= nil{
		return err
	}
	return nil
}

