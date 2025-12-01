package main

import (
	"ChatApp/internal/config"
	"ChatApp/internal/database"
)

func main() {
	cfg := config.LoadCofig()
	database.NewPostgresDB(cfg)
}
