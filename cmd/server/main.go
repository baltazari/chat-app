package main

import (
	"chat-app/internal/config"
	"chat-app/internal/database"
)

func main() {
	cfg := config.Load()
	database.Init(cfg)
}
