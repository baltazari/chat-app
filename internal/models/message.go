package models

import "time"

type Message struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"index;not null"`
	RoomID    uint   `gorm:"index;not null"`
	Content   string `gorm:"type:text;not null"`
	CreatedAt time.Time

	User User `gorm:"constraint:OnDelete:CASCADE;"`
	Room Room `gorm:"constraint:OnDelete:CASCADE;"`
}
