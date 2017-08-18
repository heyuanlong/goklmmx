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
		conn.SetReadDeadline(time.Now().Add(kother.FD_TIMEOUT_SECOND_GAME* time.Second))
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
		if msgLen == 0 {
			continue
		}
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
