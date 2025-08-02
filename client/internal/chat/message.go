package chat

import (
	"time"
)

type Message struct {
	User      string
	Content   string
	Timestamp time.Time
}
