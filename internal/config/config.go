package config

import "os"

type Config struct {
	DBurl      string
	JWTSecret  string
	ServerPort string
}

func Load() *Config {
	cfg := &Config{
		DBurl:      getENV("DATABASE_URL", "host=localhost user=mishka password=M1sh1k019891710 dbname=chat port=5432 sslmode=disable"),
		JWTSecret:  getENV("JWT_SECRET", "supersecret"),
		ServerPort: getENV("PORT", ":8080"),
	}
	return cfg
}

func getENV(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
