package core

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// ** EnvConfig struct to hold configuration values
type EnvConfig struct {
	Port      string
	DBHost    string
	DBUser    string
	DBPass    string
	DBName    string
	DBPort    int
	DBSync    bool
	JwtSecret string
}

var Config EnvConfig

func LoadEnv(file string) {
	if err := godotenv.Load(file); err != nil {
		log.Fatal("Error loading .env file")
	}

	Config = EnvConfig{
		Port:      getEnv("PORT", "3001"),
		DBHost:    getEnv("DB_HOST", "localhost"),
		DBUser:    getEnv("DB_USER", "postgres"),
		DBPass:    getEnv("DB_PASS", "postgres"),
		DBName:    getEnv("DB_NAME", "postgres"),
		DBPort:    getEnvAsInt("DB_PORT", 5432),
		DBSync:    getEnvAsBool("DB_SYNC", true),
		JwtSecret: getEnv("JWT_SECRET", "your_secret_token"),
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

func getEnvAsBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return fallback
}
