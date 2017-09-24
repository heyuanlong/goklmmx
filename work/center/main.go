package main

import (
	"net"
	klog "goklmmx/lib/log"
	kconf "goklmmx/lib/conf"
)

func main() {
	kconf.SetFile("conf/config.cfg")
	klog.SetLogfile()


	c, _ := kconf.GetConf()
	serverPort,_ := c.String("server","port")
	listen_sock,err := net.Listen("tcp",":"+serverPort)
	if err != nil{
		klog.Klog.Fatalln(err)
	}
	defer listen_sock.Close()
	for{
		_,err := listen_sock.Accept()
		if err != nil {
			klog.Klog.Println("listen_sock.Accept error:",err)
			continue
		}
		klog.Klog.Println("kkkkkkk")
	}


}