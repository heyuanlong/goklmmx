package main

import (
	"net"
	"time"
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
/*
	conn.Close() 并不会引起崩溃
	conn.Close()后，conn.Read() 会返回err
	超时，直接在read里做就ok了
*/

