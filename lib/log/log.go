package log

import (
	"fmt"
	"os"
	syslog "log"
)

var klog *syslog.Logger

func init() {
	klog = syslog.New(os.Stdout,"",syslog.LstdFlags | syslog.Lshortfile )
}

func SetLogfile(f string ) error  {
	if f != ""{
		logFile,err  := os.Create(f)
		if err != nil {
			syslog.Fatalln("open file error")
		}
		klog = syslog.New(logFile,"",syslog.LstdFlags | syslog.Lshortfile )
	}

	return nil
}

func  Common(infmt string, args ...interface{}) {
	fmts := fmt.Sprintf(infmt,args...)
	klog.Println(fmts)
}