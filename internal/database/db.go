package database

import (
	"chat-app/internal/config"
	"chat-app/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(cfg *config.Config) {
	var err error
	DB, err = gorm.Open(postgres.Open(cfg.DBurl), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to conection database", err)
	}
	if err := DB.AutoMigrate(&models.User{}, &models.Room{}, &models.Message{}); err != nil {
		log.Fatal("failed to migrate:", err)
	}
}
