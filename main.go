// Package main implements a Strava Custom Goals tracker that fetches activities
// from the Strava API and provides goal tracking functionality.
//
// Features:
// - OAuth token refresh for secure API access
// - Activity fetching with comprehensive data
// - Custom goal tracking and progress monitoring
// - Detailed activity analysis and reporting
//
// Setup:
// 1. Create a Strava API application at https://www.strava.com/settings/api
// 2. Replace the credentials below with your app's CLIENT_ID, CLIENT_SECRET, and REFRESH_TOKEN
// 3. Run: go run main.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Configuration constants - replace with your Strava API credentials
const (
	CLIENT_ID     = "your_client_id"     // Strava application client ID
	CLIENT_SECRET = "your_client_secret" // Strava application client secret
	REFRESH_TOKEN = "your_refresh_token" // OAuth refresh token
)

// API configuration
const (
	stravaTokenURL      = "https://www.strava.com/oauth/token"
	stravaActivitiesURL = "https://www.strava.com/api/v3/athlete/activities"
	defaultPerPage      = 30
	requestTimeout      = 30 * time.Second
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

// StravaClient handles API interactions with Strava
type StravaClient struct {
	ClientID     string
	ClientSecret string
	RefreshToken string
	httpClient   *http.Client
}

// NewStravaClient creates a new Strava API client
func NewStravaClient(clientID, clientSecret, refreshToken string) *StravaClient {
	return &StravaClient{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RefreshToken: refreshToken,
		httpClient:   &http.Client{Timeout: requestTimeout},
	}
}

func main() {
	log.Println("🚀 Strava Custom Goals Tracker Starting...")

	// Initialize Strava client
	client := NewStravaClient(CLIENT_ID, CLIENT_SECRET, REFRESH_TOKEN)

	// Get access token
	log.Println("📡 Authenticating with Strava API...")
	accessToken, err := client.GetAccessToken()
	if err != nil {
		log.Fatalf("❌ Authentication failed: %v", err)
	}
	log.Println("✅ Successfully authenticated")

	// Fetch recent activities
	log.Println("📊 Fetching recent activities...")
	activities, err := client.GetActivities(accessToken)
	if err != nil {
		log.Fatalf("❌ Failed to fetch activities: %v", err)
	}
	log.Printf("✅ Retrieved %d activities", len(activities))

	// Process and display activities
	if len(activities) == 0 {
		log.Println("ℹ️ No activities found")
		return
	}

	// Enhance activities with calculated fields
	for i := range activities {
		activities[i].enhanceWithCalculatedFields()
	}

	// Display activity summary
	displayActivities(activities)

	log.Printf("🎯 Analysis complete: processed %d activities", len(activities))
}

// GetAccessToken exchanges refresh token for access token via OAuth
func (c *StravaClient) GetAccessToken() (string, error) {
	requestBody := map[string]string{
		"client_id":     c.ClientID,
		"client_secret": c.ClientSecret,
		"refresh_token": c.RefreshToken,
		"grant_type":    "refresh_token",
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("marshal request: %w", err)
	}

	resp, err := c.httpClient.Post(stravaTokenURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("token request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("token API error %d: %s", resp.StatusCode, string(body))
	}

	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", fmt.Errorf("decode token response: %w", err)
	}

	if tokenResp.AccessToken == "" {
		return "", fmt.Errorf("empty access token received")
	}

	return tokenResp.AccessToken, nil
}

// GetActivities fetches recent activities using the provided access token
func (c *StravaClient) GetActivities(accessToken string) ([]Activity, error) {
	url := fmt.Sprintf("%s?per_page=%d&page=1", stravaActivitiesURL, defaultPerPage)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("activities API error %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	var activities []Activity
	if err := json.Unmarshal(body, &activities); err != nil {
		return nil, fmt.Errorf("unmarshal activities: %w", err)
	}

	return activities, nil
}

// enhanceWithCalculatedFields adds calculated fields to an activity
func (a *Activity) enhanceWithCalculatedFields() {
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

// displayActivities shows activity information in a formatted way
func displayActivities(activities []Activity) {
	fmt.Println("\n🏃‍♂️ === RECENT ACTIVITIES ===")

	for i, activity := range activities {
		fmt.Printf("\n📈 Activity %d\n", i+1)
		fmt.Printf("   🏷️  Name: %s\n", activity.Name)
		fmt.Printf("   🎯 Type: %s\n", activity.Type)
		fmt.Printf("   📏 Distance: %.2f km\n", activity.DistanceKm)
		fmt.Printf("   ⏱️  Moving Time: %s\n", formatDuration(activity.MovingTime))

		if activity.TotalElevGain > 0 {
			fmt.Printf("   ⛰️  Elevation Gain: %.0f m\n", activity.TotalElevGain)
		}

		if activity.Type == "Run" && activity.PaceMinPerKm != "" {
			fmt.Printf("   🏃 Average Pace: %s min/km\n", activity.PaceMinPerKm)
		}

		if activity.HasHeartrate && activity.AverageHeartrate > 0 {
			fmt.Printf("   ❤️  Avg Heart Rate: %.0f bpm\n", activity.AverageHeartrate)
		}

		if activity.Kudos > 0 {
			fmt.Printf("   👍 Kudos: %d\n", activity.Kudos)
		}

		fmt.Printf("   📅 Date: %s\n", formatDate(activity.StartDateLocal))
	}
}

// formatDuration converts seconds to human-readable duration format
func formatDuration(seconds int) string {
	duration := time.Duration(seconds) * time.Second
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	secs := int(duration.Seconds()) % 60

	if hours > 0 {
		return fmt.Sprintf("%dh %dm %ds", hours, minutes, secs)
	}
	return fmt.Sprintf("%dm %ds", minutes, secs)
}

// formatDate converts ISO 8601 date string to readable format
func formatDate(dateStr string) string {
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
