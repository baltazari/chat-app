package models

import "time"

type Room struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"uniqueIndex;size:150;not null"`
	CreateAt time.Time
	UpdateAt time.Time
	Messages []Message
}
