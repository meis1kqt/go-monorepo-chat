package domain

import "time"


type Dialog struct {
	ID int64
	CreatorID int64
	UserID int64
	CreatedAt time.Time
}