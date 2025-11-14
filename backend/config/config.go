package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	DatabaseURL string
	JWTSecret   string
	Port        string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://notesuser:notespass@localhost:5432/notes_db?sslmode=disable"),
		JWTSecret:   getEnv("JWT_SECRET", "mysupersecretkey123"),
		Port:        getEnv("PORT", "8080"),
	}
}

// getEnv gets environment variable with a default fallback
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}