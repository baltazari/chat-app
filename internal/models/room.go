package models

import "time"

type Room struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"uniqueIndex;size:150;not null"`
	CreatedAt time.Time
	Updatedt  time.Time
	Messages  []Message
}
