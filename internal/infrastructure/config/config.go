package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	ServerPort    string
	DatabaseURL   string
	JWTSecret     string
	JWTExpiration time.Duration
}

func Load() *Config {
	return &Config{
		ServerPort:    getEnv("SERVER_PORT", ":8080"),
		DatabaseURL:   getEnv("DATABASE_URL", "postgresql://user:password@localhost/real_estate"),
		JWTSecret:     getEnv("JWT_SECRET", "super-secret-key-change-in-production"),
		JWTExpiration: getEnvDuration("JWT_EXPIRATION", 24) * time.Hour,
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvDuration(key string, defaultValue int) time.Duration {
	if value := os.Getenv(key); value != "" {
		if hours, err := strconv.Atoi(value); err == nil {
			return time.Duration(hours)
		}
	}
	return time.Duration(defaultValue)
}
