package models

import "time"

type Massage struct {
	ID        uint
	UserID    uint
	RoomID    uint
	User      User
	Room      Room
	Content   string
	CreatedAt time.Time
}
