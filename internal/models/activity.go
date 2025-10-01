// Package models contains data structures for Strava API responses and activity data.
package models

import (
	"fmt"
	"time"
)

// TokenResponse represents the OAuth token response from Strava API
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
	TokenType    string `json:"token_type"`
}

// Activity represents a Strava activity with enhanced tracking data
type Activity struct {
	ID               int64   `json:"id"`
	Name             string  `json:"name"`
	Distance         float64 `json:"distance"`             // meters
	MovingTime       int     `json:"moving_time"`          // seconds
	ElapsedTime      int     `json:"elapsed_time"`         // seconds
	TotalElevGain    float64 `json:"total_elevation_gain"` // meters
	Type             string  `json:"type"`                 // activity type (Run, Ride, etc.)
	StartDate        string  `json:"start_date"`           // ISO 8601 format
	StartDateLocal   string  `json:"start_date_local"`     // local timezone
	AverageSpeed     float64 `json:"average_speed"`        // m/s
	MaxSpeed         float64 `json:"max_speed"`            // m/s
	HasHeartrate     bool    `json:"has_heartrate"`
	AverageHeartrate float64 `json:"average_heartrate"` // bpm
	Kudos            int     `json:"kudos_count"`

	// Calculated fields for enhanced analysis
	DistanceKm      float64 `json:"-"`
	MovingTimeHours float64 `json:"-"`
	PaceMinPerKm    string  `json:"-"`
}

// EnhanceWithCalculatedFields adds calculated fields to an activity
func (a *Activity) EnhanceWithCalculatedFields() {
	// Convert distances and times to more readable units
	a.DistanceKm = a.Distance / 1000
	a.MovingTimeHours = float64(a.MovingTime) / 3600

	// Calculate pace for running activities
	if a.Type == "Run" && a.Distance > 0 {
		paceSecondsPerKm := float64(a.MovingTime) / (a.Distance / 1000)
		minutes := int(paceSecondsPerKm / 60)
		seconds := int(paceSecondsPerKm) % 60
		a.PaceMinPerKm = fmt.Sprintf("%d:%02d", minutes, seconds)
	}
}

// FormatDuration converts seconds to human-readable duration format
func FormatDuration(seconds int) string {
	duration := time.Duration(seconds) * time.Second
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	secs := int(duration.Seconds()) % 60

	if hours > 0 {
		return fmt.Sprintf("%dh %dm %ds", hours, minutes, secs)
	}
	return fmt.Sprintf("%dm %ds", minutes, secs)
}

// FormatDate converts ISO 8601 date string to readable format
func FormatDate(dateStr string) string {
	if dateStr == "" {
		return "N/A"
	}

	// Parse the ISO 8601 format
	t, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return dateStr // Return original if parsing fails
	}

	// Format to a more readable format
	return t.Format("Jan 2, 2006 15:04")
}
