package database

import (
	"ChatApp/internal/config"
	"ChatApp/internal/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewPostgresDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPasswod,
		cfg.DBName,
		cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("faild  connect to to database: %v", err)
	}
	if err := db.AutoMigrate(&models.User{}, &models.Room{}, &models.Message{}); err != nil {
		log.Fatalf("faild to migrate database: %v", err)
	}
	return db
}
