package main

import (
	"net"
	"sync"
	"time"
	"errors"
)

var userConnMap 	map[int] net.Conn
var rwm			 	*sync.RWMutex

func init() {
	userConnMap = make(map[int] net.Conn)
	rwm = new(sync.RWMutex)
}

func AddUserConnMap(id int , conn net.Conn) error  {
	rwm.Lock()
	userConnMap[id] = conn
	rwm.Unlock()
	return nil
}
func DelUserConnMap(id int) error  {
	rwm.Lock()
	detele(userConnMap,id)
	rwm.Unlock()
	return nil
}

func GetUserConnMap(id int) (net.Conn, error)  {
	rwm.RLock()
	v, ok := userConnMap[id]
	rwm.RUnlock()

	if ok {
		return v,nil
	} else {
		return nil,errors.New("not id connect")
	}
}