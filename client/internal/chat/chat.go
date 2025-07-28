package chat

import (
	"time"
)

type Message struct {
	User      string
	Content   string
	Timestamp time.Time
}

type State struct {
	User       string
	Status     ConnStatus
	History    []Message
	newMessage chan Message
}
