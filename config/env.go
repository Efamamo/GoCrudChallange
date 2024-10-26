package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application's configuration values.
type Config struct {
	Port string
	Host string
}

// Envs holds the application's configuration loaded from environment variables.
var Envs = initializeConfig()

func initializeConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, proceeding with defaults.")
	}

	return Config{
		Host: getEnv("HOST", "localhost"),
		Port: getEnv("PORT", "8080"),
	}
}

// getEnv retrieves the value of an environment variable or returns a fallback value if the variable is not set.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
