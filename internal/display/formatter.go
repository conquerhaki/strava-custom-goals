// Package display provides formatting and display functionality for Strava activities.
package display

import (
	"fmt"

	"strava-custom-goals/internal/goals"
	"strava-custom-goals/internal/models"
)

// DisplayActivities shows activity information in a formatted way
func DisplayActivities(activities []models.Activity) {
	fmt.Println("\n🏃‍♂️ === RECENT ACTIVITIES ===")

	for i, activity := range activities {
		fmt.Printf("\n📈 Activity %d\n", i+1)
		fmt.Printf("   🏷️  Name: %s\n", activity.Name)
		fmt.Printf("   🎯 Type: %s\n", activity.Type)
		fmt.Printf("   📏 Distance: %.2f km\n", activity.DistanceKm)
		fmt.Printf("   ⏱️  Moving Time: %s\n", models.FormatDuration(activity.MovingTime))

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

		fmt.Printf("   📅 Date: %s\n", models.FormatDate(activity.StartDateLocal))
	}
}

// DisplaySummary shows a summary of activities
func DisplaySummary(activities []models.Activity) {
	if len(activities) == 0 {
		fmt.Println("📊 No activities to analyze")
		return
	}

	totalDistance := 0.0
	totalTime := 0
	runCount := 0
	rideCount := 0

	for _, activity := range activities {
		totalDistance += activity.DistanceKm
		totalTime += activity.MovingTime

		switch activity.Type {
		case "Run":
			runCount++
		case "Ride":
			rideCount++
		}
	}

	fmt.Println("\n📊 === ACTIVITY SUMMARY ===")
	fmt.Printf("   📈 Total Activities: %d\n", len(activities))
	fmt.Printf("   🏃 Runs: %d\n", runCount)
	fmt.Printf("   🚴 Rides: %d\n", rideCount)
	fmt.Printf("   📏 Total Distance: %.2f km\n", totalDistance)
	fmt.Printf("   ⏱️  Total Time: %s\n", models.FormatDuration(totalTime))

	if totalDistance > 0 {
		avgDistance := totalDistance / float64(len(activities))
		fmt.Printf("   📊 Average Distance: %.2f km\n", avgDistance)
	}
}

// DisplayWeeklyGoalsProgress shows progress toward weekly fitness goals
func DisplayWeeklyGoalsProgress(progress *goals.WeeklyProgress) {
	fmt.Println("\n🎯 === WEEKLY GOALS PROGRESS ===")
	
	// Running progress
	runningPercent := progress.GetRunningProgressPercentage()
	runningStatus := "🔴"
	if progress.IsRunningGoalAchieved() {
		runningStatus = "✅"
	} else if runningPercent >= 75 {
		runningStatus = "🟡"
	} else if runningPercent >= 50 {
		runningStatus = "🟠"
	}
	
	fmt.Printf("   🏃‍♂️ Running Goal: %.1f/%.1f km %s (%.1f%%)\n", 
		progress.RunningDistance, progress.Goals.RunningGoalKm, runningStatus, runningPercent)
	
	if !progress.IsRunningGoalAchieved() {
		fmt.Printf("      💭 Need %.1f km more to reach your goal\n", progress.GetRunningRemainingDistance())
	}
	
	// Workout progress
	workoutPercent := progress.GetWorkoutProgressPercentage()
	workoutStatus := "🔴"
	if progress.IsWorkoutGoalAchieved() {
		workoutStatus = "✅"
	} else if workoutPercent >= 75 {
		workoutStatus = "🟡"
	} else if workoutPercent >= 50 {
		workoutStatus = "🟠"
	}
	
	fmt.Printf("   💪 Workout Goal: %.1f/%.1f hours %s (%.1f%%)\n", 
		progress.WorkoutHours, progress.Goals.WorkoutGoalHours, workoutStatus, workoutPercent)
	
	if !progress.IsWorkoutGoalAchieved() {
		fmt.Printf("      💭 Need %.1f hours more to reach your goal\n", progress.GetWorkoutRemainingHours())
	}
	
	// Activity counts for this week
	fmt.Printf("   📊 This Week: %d runs, %d workouts, %d total activities\n", 
		progress.RunCount, progress.WorkoutCount, progress.TotalActivities)
	
	// Motivational message
	fmt.Printf("\n   %s\n", progress.GetMotivationalMessage())
}
