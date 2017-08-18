package main

import (
	"net"

	klog "goklmmx/lib/log"
	kconf "goklmmx/lib/conf"
	kredis "goklmmx/lib/db/redis"
	kmysql "goklmmx/lib/db/mysql"
)


func main() {
	kconf.SetFile("conf/config.cfg")
	klog.SetLogfile()
	kredis.RedisInit()
	kmysql.MysqlInit()

	kmysql.Test()

	c, _ := kconf.GetConf()
	serverPort,_ := c.String("server","port")
	listen_sock,err := net.Listen("tcp",":"+serverPort)
	if err != nil{
		klog.Klog.Fatalln(err)
	}
	defer listen_sock.Close()
	for{
		new_conn,err := listen_sock.Accept()
		if err != nil {
			klog.Klog.Println("listen_sock.Accept error:",err)
			continue
		}
		go HandleClient(new_conn)
	}
}