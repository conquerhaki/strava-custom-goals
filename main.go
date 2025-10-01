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
// 2. Copy .env.example to .env and fill in your credentials
// 3. Run: go run main.go
package main

import (
	"flag"
	"log"

	"strava-custom-goals/config"
	"strava-custom-goals/internal/client"
	"strava-custom-goals/internal/display"
	"strava-custom-goals/internal/goals"
)

func main() {
	// Parse command line flags
	var (
		showHelp    = flag.Bool("help", false, "Show help message")
		maxResults  = flag.Int("max", 30, "Maximum number of activities to fetch")
		showSummary = flag.Bool("summary", true, "Show activity summary")
		showDetails = flag.Bool("details", true, "Show detailed activities")
	)
	flag.Parse()

	if *showHelp {
		flag.Usage()
		return
	}

	log.Println("🚀 Strava Custom Goals Tracker Starting...")

	// Load configuration from environment variables
	cfg := config.LoadConfig()

	// Initialize Strava client
	stravaClient := client.NewStravaClient(cfg.ClientID, cfg.ClientSecret, cfg.RefreshToken)

	// Get access token
	log.Println("📡 Authenticating with Strava API...")
	accessToken, err := stravaClient.GetAccessToken()
	if err != nil {
		log.Fatalf("❌ Authentication failed: %v", err)
	}
	log.Println("✅ Successfully authenticated")

	// Fetch recent activities
	log.Println("📊 Fetching recent activities...")
	activities, err := stravaClient.GetActivities(accessToken)
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
		activities[i].EnhanceWithCalculatedFields()
	}

	// Calculate weekly goals progress
	log.Println("🎯 Calculating weekly goals progress...")
	weeklyGoals := goals.WeeklyGoals{
		RunningGoalKm:    cfg.WeeklyRunningGoalKm,
		WorkoutGoalHours: cfg.WeeklyWorkoutGoalHours,
	}
	weeklyProgress := goals.CalculateWeeklyProgress(activities, weeklyGoals)

	// Display weekly goals progress
	display.DisplayWeeklyGoalsProgress(weeklyProgress)

	// Display detailed activities (if requested)
	if *showDetails {
		display.DisplayActivities(activities[:min(len(activities), *maxResults)])
	}

	// Display summary (if requested)
	if *showSummary {
		display.DisplaySummary(activities)
	}

	log.Printf("🎯 Analysis complete: processed %d activities", len(activities))
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
