// Package config contains configuration constants and settings for the Strava Custom Goals application.
package config

import "time"

// Strava API credentials - replace with your own values
const (
	ClientID     = "client-id"     // Strava application client ID
	ClientSecret = "client-secret" // Strava application client secret
	RefreshToken = "refresh-token" // OAuth refresh token
)

// API endpoints and configuration
const (
	StravaTokenURL      = "https://www.strava.com/oauth/token"
	StravaActivitiesURL = "https://www.strava.com/api/v3/athlete/activities"
	DefaultPerPage      = 30
	RequestTimeout      = 30 * time.Second
)
