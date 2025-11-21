package models

import "time"

type Message struct {
	ID        uint
	UserID    uint
	RoomID    uint
	User      User
	Room      Room
	Content   string
	CreatedAt time.Time
}
