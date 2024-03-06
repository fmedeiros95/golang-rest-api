package core

import (
	"os"
	"strconv"
)

// EnvConfig struct to hold configuration values
type EnvConfig struct {
	Port      string
	DBHost    string
	DBUser    string
	DBPass    string
	DBName    string
	DBPort    int
	JwtSecret string
}

var Config EnvConfig

func LoadEnv() {
	Config = EnvConfig{
		Port:      getEnv("PORT", "8080"),
		DBHost:    getEnv("DB_HOST", "localhost"),
		DBUser:    getEnv("DB_USER", "postgres"),
		DBPass:    getEnv("DB_PASS", "postgres"),
		DBName:    getEnv("DB_NAME", "postgres"),
		DBPort:    getEnvAsInt("DB_PORT", 5432),
		JwtSecret: getEnv("JWT_SECRET", "your_secret_key"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return fallback
}
