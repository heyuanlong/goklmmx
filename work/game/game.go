package main

import (

	"net"
	"time"


	klog "goklmmx/lib/log"
	knet "goklmmx/lib/net"
	kother "goklmmx/lib/other"
)

const (
	G_MSG_SIZE_MAX    = 1024
	G_BUF_SIZE_MAX    = 65536
)

func HandleClient(conn net.Conn)  {
	klog.Klog.Println("HandleClient")
	defer conn.Close()
	agent := NewAgent(conn)

	var bufBuf = make([]byte,0)
	var msgBuf = make([]byte, G_MSG_SIZE_MAX)
	for  {
		conn.SetReadDeadline(time.Now().Add(time.Duration(kother.FD_TIMEOUT_SECOND_GAME) * time.Second))
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

		klog.Klog.Println("msgType:",msgType)
		err = nil
		err = agent.DealPackage(msgType,pBuf)
		if err != nil{
			klog.Klog.Println(err)
			return
		}
		bufBuf = bufBuf[msgLen:]
	}
}
