package dealChan1

import (
	kbase "goklmmx/work/game/base"
	klog "goklmmx/lib/log"
)
func Deal()  {
	for  {
		ms := <- kbase.Chan1
		klog.Klog.Println("chan1 ",ms.MsgType)
	}
}