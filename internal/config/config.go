package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	AppEnv    string
	AppPort   string
	DBHost    string
	DBPort    int
	DBUser    string
	DBPasswod string
	DBName    string

	JWTSecret       string
	JWTExpirePeriod time.Duration
}

func LoadCofig() *Config {
	_ = godotenv.Load() // ignor error if .env not found
	dbport, err := strconv.Atoi(getEnv("DBPORT", "5432"))
	if err != nil {
		log.Fatalf("invalid DB_PORT: %v", err)
	}
	jwtMinutes, err := strconv.Atoi(getEnv("JWT_EXPIRE_MINUTES", "1440"))
	if err != nil {
		log.Fatalf("invalid JWT_EXPIRE_MINUTES: %v", err)
	}

	return &Config{
		AppEnv:  getEnv("APP_ENV", "development"),
		AppPort: getEnv("APP_PORT", "8080"),

		DBHost:    getEnv("DB_HOST", "localhost"),
		DBPort:    dbport,
		DBUser:    getEnv("DB_USER", "mishka"),
		DBPasswod: getEnv("DB_PASSWORD", "M1sh1k019891710"),
		DBName:    getEnv("DB_NAME", "chat"),

		JWTSecret:       getEnv("JWT_SECRET", "change_me"),
		JWTExpirePeriod: time.Duration(jwtMinutes) * time.Minute,
	}
}

func getEnv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return def
}
