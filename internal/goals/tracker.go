// Package goals provides functionality for tracking weekly fitness goals
package goals

import (
	"time"

	"strava-custom-goals/internal/models"
)

// WeeklyGoals represents weekly fitness targets
type WeeklyGoals struct {
	RunningGoalKm    float64
	WorkoutGoalHours float64
}

// WeeklyProgress tracks progress toward weekly goals
type WeeklyProgress struct {
	Goals           WeeklyGoals
	RunningDistance float64
	WorkoutHours    float64
	TotalActivities int
	RunCount        int
	WorkoutCount    int
}

// CalculateWeeklyProgress calculates progress toward weekly goals from activities
func CalculateWeeklyProgress(activities []models.Activity, goals WeeklyGoals) *WeeklyProgress {
	progress := &WeeklyProgress{
		Goals: goals,
	}

	// Get the start of current week (Monday)
	now := time.Now()
	weekday := now.Weekday()
	if weekday == time.Sunday {
		weekday = 7 // Treat Sunday as day 7
	}
	weekStart := now.AddDate(0, 0, -int(weekday-time.Monday))
	weekStart = time.Date(weekStart.Year(), weekStart.Month(), weekStart.Day(), 0, 0, 0, 0, weekStart.Location())

	for _, activity := range activities {
		// Parse activity start date
		activityTime, err := time.Parse(time.RFC3339, activity.StartDate)
		if err != nil {
			continue // Skip activities with invalid dates
		}

		// Only count activities from current week
		if activityTime.Before(weekStart) {
			continue
		}

		progress.TotalActivities++

		// Track running activities
		if activity.Type == "Run" {
			progress.RunningDistance += activity.DistanceKm
			progress.RunCount++
		}

		// Track workout activities (gym, weight lifting, strength training)
		if isWorkoutActivity(activity.Type) {
			progress.WorkoutHours += activity.MovingTimeHours
			progress.WorkoutCount++
		}
	}

	return progress
}

// isWorkoutActivity determines if an activity type counts as a workout
func isWorkoutActivity(activityType string) bool {
	workoutTypes := map[string]bool{
		"WeightTraining":  true,
		"Workout":         true,
		"Crossfit":        true,
		"StairStepper":    true,
		"Elliptical":      true,
		"Yoga":           true,
		"Pilates":        true,
		"RockClimbing":   true,
		"Swimming":       true,
	}
	return workoutTypes[activityType]
}

// GetRunningProgressPercentage returns running progress as percentage
func (p *WeeklyProgress) GetRunningProgressPercentage() float64 {
	if p.Goals.RunningGoalKm == 0 {
		return 0
	}
	return (p.RunningDistance / p.Goals.RunningGoalKm) * 100
}

// GetWorkoutProgressPercentage returns workout progress as percentage
func (p *WeeklyProgress) GetWorkoutProgressPercentage() float64 {
	if p.Goals.WorkoutGoalHours == 0 {
		return 0
	}
	return (p.WorkoutHours / p.Goals.WorkoutGoalHours) * 100
}

// IsRunningGoalAchieved checks if running goal is met
func (p *WeeklyProgress) IsRunningGoalAchieved() bool {
	return p.RunningDistance >= p.Goals.RunningGoalKm
}

// IsWorkoutGoalAchieved checks if workout goal is met
func (p *WeeklyProgress) IsWorkoutGoalAchieved() bool {
	return p.WorkoutHours >= p.Goals.WorkoutGoalHours
}

// GetRunningRemainingDistance returns remaining distance to achieve running goal
func (p *WeeklyProgress) GetRunningRemainingDistance() float64 {
	remaining := p.Goals.RunningGoalKm - p.RunningDistance
	if remaining < 0 {
		return 0
	}
	return remaining
}

// GetWorkoutRemainingHours returns remaining hours to achieve workout goal
func (p *WeeklyProgress) GetWorkoutRemainingHours() float64 {
	remaining := p.Goals.WorkoutGoalHours - p.WorkoutHours
	if remaining < 0 {
		return 0
	}
	return remaining
}

// GetMotivationalMessage returns a motivational message based on progress
func (p *WeeklyProgress) GetMotivationalMessage() string {
	runningAchieved := p.IsRunningGoalAchieved()
	workoutAchieved := p.IsWorkoutGoalAchieved()

	if runningAchieved && workoutAchieved {
		return "🎉 Congratulations! You've achieved both your running and workout goals this week!"
	} else if runningAchieved {
		return "🏃‍♂️ Great job on your running goal! Keep up the momentum with your workouts!"
	} else if workoutAchieved {
		return "💪 Excellent work on your workout goal! Time to lace up those running shoes!"
	} else {
		runningPercent := p.GetRunningProgressPercentage()
		workoutPercent := p.GetWorkoutProgressPercentage()
		
		if runningPercent > 50 && workoutPercent > 50 {
			return "🔥 You're over halfway to both goals! Keep pushing!"
		} else if runningPercent > workoutPercent {
			return "🏃‍♂️ Strong running progress! Time to balance it with some strength training!"
		} else if workoutPercent > runningPercent {
			return "💪 Great workout momentum! Add some cardio to complete the balance!"
		} else {
			return "🚀 The week is young! Time to start building towards your goals!"
		}
	}
}
