package net

import (
	"goklmmx/lib/utils"
)

//   2 + 4 + n
func ParsePackage(buf []byte) (msgLen int,msgType int,pBuf []byte) {
	zlen := len(buf)
	if zlen < 6{
		return 0,0,nil
	}
	msgLen = int(utils.BytesToUint16(buf[0:2]))
	if zlen < msgLen {
		return 0,0,nil
	}
	msgType = utils.BytesToInt(buf[2:6])
	pBuf = buf[6:msgLen]
	return
}