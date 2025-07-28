package chat

import (
	"fmt"
	"net"
)

type ConnStatus uint64

const (
	Unknown ConnStatus = iota
	Online
	Offline
)

func (s ConnStatus) String() string {
	switch s {
	case Online:
		return "Online"
	case Offline:
		return "Offline"
	case Unknown:
		return "Unknown"
	default:
		return fmt.Sprintf("ConnStatus-(%d)", s)
	}
}

type ConnHandler struct {
	Socket string
	Conn   net.Conn
	Status ConnStatus
}
