package chat

import (
	"net"
)

type ConnHandler struct {
	Socket  string
	Conn    net.Conn
	Status  ConnStatus
	Recieve chan Message
}
