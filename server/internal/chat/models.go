package chat

import (
	"time"
)

type User struct {
	Username string
	Name     string
}

type Message struct {
	Content   string
	Sender    *User
	CreatedAt time.Time
}
