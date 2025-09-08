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

func Load() Config {
    _ = godotenv.Load()
    cfg := Config{
        DatabaseURL: os.Getenv("DATABASE_URL"),
        HTTPAddr:    getenv("HTTP_ADDR", ":8080"),
    }
    if cfg.DatabaseURL == "" {
        log.Println("warning: DATABASE_URL is empty; use in-memory repo or set env")
    }
    return cfg
}

func getenv(key, def string) string {
    if v := os.Getenv(key); v != "" { return v }
    return def
}


