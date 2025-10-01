// Package config contains configuration constants and settings for the Strava Custom Goals application.
package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Config holds all configuration values
type Config struct {
	ClientID     string
	ClientSecret string
	RefreshToken string
}

// API endpoints and configuration constants
const (
	StravaTokenURL      = "https://www.strava.com/oauth/token"
	StravaActivitiesURL = "https://www.strava.com/api/v3/athlete/activities"
	DefaultPerPage      = 30
	RequestTimeout      = 30 * time.Second
)

// LoadConfig loads configuration from environment variables and .env file
func LoadConfig() *Config {
	// Load .env file if it exists (ignore errors for production deployments)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	config := &Config{
		ClientID:     getEnvOrDefault("STRAVA_CLIENT_ID", ""),
		ClientSecret: getEnvOrDefault("STRAVA_CLIENT_SECRET", ""),
		RefreshToken: getEnvOrDefault("STRAVA_REFRESH_TOKEN", ""),
	}

	// Validate required configuration
	if config.ClientID == "" || config.ClientSecret == "" || config.RefreshToken == "" {
		log.Fatal("❌ Missing required environment variables. Please check your .env file or set STRAVA_CLIENT_ID, STRAVA_CLIENT_SECRET, and STRAVA_REFRESH_TOKEN")
	}

	return config
}

// getEnvOrDefault gets an environment variable or returns a default value
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
