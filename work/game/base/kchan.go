package base

type MsgChan struct {
	MsgType int
	Msg []byte
}

var Chan1 chan MsgChan

func init() {
	Chan1 = make(chan MsgChan,1000)
}