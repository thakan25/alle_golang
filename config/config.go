package config

import (
	"os"
	"strconv"
)

// Config holds all application configuration
type Config struct {
	ServerPort     int
	LogLevel       string
	EnableCORS     bool
	MaxRequestSize int64
}

// LoadConfig loads configuration from environment variables with defaults
func LoadConfig() *Config {
	return &Config{
		ServerPort:     getEnvInt("PORT", 8080),
		LogLevel:       getEnvString("LOG_LEVEL", "info"),
		EnableCORS:     getEnvBool("ENABLE_CORS", true),
		MaxRequestSize: getEnvInt64("MAX_REQUEST_SIZE", 1048576), // 1MB default
	}
}

// Helper functions to get environment variables with defaults
func getEnvString(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}

func getEnvInt64(key string, defaultValue int64) int64 {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
			return intValue
		}
	}
	return defaultValue
} 