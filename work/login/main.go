package main

import (

	klog "goklmmx/lib/log"
	kconf "goklmmx/lib/conf"
)


func main() {
	kconf.SetFile("conf/config.cfg")
	c, _ := kconf.GetConf()
	logfile,_ := c.String("server","logfile")
	klog.SetLogfile(logfile)



	klog.Common("kkkkkkkkkkkkkkkkkkkkkk:%d",12)
}