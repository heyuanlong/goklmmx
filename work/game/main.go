package main

import (
	"net"
	klog "goklmmx/lib/log"
	kconf "goklmmx/lib/conf"
	kredis "goklmmx/lib/db/redis"
	kmysql "goklmmx/lib/db/mysql"
	kagentManager "goklmmx/work/game/agentManager"
	kdealChan1 "goklmmx/work/game/dealChan1"
)


func main() {
	kconf.SetFile("conf/config.cfg")
	klog.SetLogfile()
	kredis.RedisInit()
	kmysql.MysqlInit()

	go kdealChan1.Deal()

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
		go kagentManager.HandleClient(new_conn)
	}
}
/*
	conn.Close() 并不会引起崩溃
	conn.Close()后，conn.Read() 会返回err
	超时，直接在read里做就ok了
*/

