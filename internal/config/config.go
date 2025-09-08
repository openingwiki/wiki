package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	HTTPAddr    string
}

func Load() *Config {
	loadDotEnv()

	cfg := &Config{
		DatabaseURL: getEnv("DATABASE_URL", ""),
		HTTPAddr:    getEnv("HTTP_ADDR", ":8080"),
	}

	if cfg.DatabaseURL == "" {
		log.Println("warning: DATABASE_URL is empty; using in-memory repository or set env")
	}

	return cfg
}

func loadDotEnv() {
	if err := godotenv.Load(); err != nil {
		// No .env file is fine, just skip
		log.Println("info: no .env file found, relying on environment variables")
	}
}

func getEnv(key, def string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return def
}
