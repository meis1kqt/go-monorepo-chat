package domain

import "time"

type Message struct {
	ID int64
	DialogID int64
	From int64
	Text string
	CreatedAt time.Time
}

