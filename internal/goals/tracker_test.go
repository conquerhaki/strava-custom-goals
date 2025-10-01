package goals

import (
	"testing"
	"time"

	"strava-custom-goals/internal/models"
)

func TestCalculateWeeklyProgress(t *testing.T) {
	// Test data - create mock activities
	now := time.Now()
	weekStart := now.AddDate(0, 0, -int(now.Weekday()-time.Monday))

	activities := []models.Activity{
		{
			Type:       "Run",
			Distance:   5000, // 5km in meters
			MovingTime: 1800, // 30 minutes
			StartDate:  weekStart.Add(24 * time.Hour).Format(time.RFC3339),
			DistanceKm: 5.0,
		},
		{
			Type:            "WeightTraining",
			MovingTime:      3600, // 1 hour
			StartDate:       weekStart.Add(48 * time.Hour).Format(time.RFC3339),
			MovingTimeHours: 1.0,
		},
		{
			Type:       "Run",
			Distance:   3000,                                             // 3km
			MovingTime: 1200,                                             // 20 minutes
			StartDate:  weekStart.AddDate(0, 0, -7).Format(time.RFC3339), // Last week - should be ignored
			DistanceKm: 3.0,
		},
	}

	goals := WeeklyGoals{
		RunningGoalKm:    10.0,
		WorkoutGoalHours: 2.0,
	}

	progress := CalculateWeeklyProgress(activities, goals)

	// Test running progress
	if progress.RunningDistance != 5.0 {
		t.Errorf("Expected running distance 5.0, got %f", progress.RunningDistance)
	}

	// Test workout progress
	if progress.WorkoutHours != 1.0 {
		t.Errorf("Expected workout hours 1.0, got %f", progress.WorkoutHours)
	}

	// Test counts
	if progress.RunCount != 1 {
		t.Errorf("Expected run count 1, got %d", progress.RunCount)
	}

	if progress.WorkoutCount != 1 {
		t.Errorf("Expected workout count 1, got %d", progress.WorkoutCount)
	}

	// Test total activities (should only count current week)
	if progress.TotalActivities != 2 {
		t.Errorf("Expected total activities 2, got %d", progress.TotalActivities)
	}
}

func TestGetRunningProgressPercentage(t *testing.T) {
	progress := &WeeklyProgress{
		Goals:           WeeklyGoals{RunningGoalKm: 10.0},
		RunningDistance: 7.5,
	}

	expected := 75.0
	actual := progress.GetRunningProgressPercentage()

	if actual != expected {
		t.Errorf("Expected percentage %f, got %f", expected, actual)
	}
}

func TestIsWorkoutActivity(t *testing.T) {
	testCases := []struct {
		activityType string
		expected     bool
	}{
		{"WeightTraining", true},
		{"Workout", true},
		{"Crossfit", true},
		{"Run", false},
		{"Walk", false},
		{"Ride", false},
		{"Swimming", true},
	}

	for _, tc := range testCases {
		result := isWorkoutActivity(tc.activityType)
		if result != tc.expected {
			t.Errorf("For activity type %s, expected %v, got %v", tc.activityType, tc.expected, result)
		}
	}
}
