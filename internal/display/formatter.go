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
	runningStatus, runningBar := getProgressDisplay(runningPercent)
	
	fmt.Printf("   🏃‍♂️ Running Target: %.1f km / %.1f km (%.1f%%)\n", 
		progress.RunningDistance, progress.Goals.RunningGoalKm, runningPercent)
	fmt.Printf("      %s %s\n", runningBar, runningStatus)
	
	if !progress.IsRunningGoalAchieved() {
		fmt.Printf("      💭 Still need: %.1f km to complete your weekly goal\n", progress.GetRunningRemainingDistance())
	} else {
		fmt.Printf("      🎉 Goal achieved! You've exceeded by %.1f km\n", progress.RunningDistance-progress.Goals.RunningGoalKm)
	}
	
	// Workout progress
	workoutPercent := progress.GetWorkoutProgressPercentage()
	workoutStatus, workoutBar := getProgressDisplay(workoutPercent)
	
	fmt.Printf("\n   💪 Workout Target: %.1f hours / %.1f hours (%.1f%%)\n", 
		progress.WorkoutHours, progress.Goals.WorkoutGoalHours, workoutPercent)
	fmt.Printf("      %s %s\n", workoutBar, workoutStatus)
	
	if !progress.IsWorkoutGoalAchieved() {
		workoutRemaining := progress.GetWorkoutRemainingHours()
		hours := int(workoutRemaining)
		minutes := int((workoutRemaining - float64(hours)) * 60)
		if hours > 0 {
			fmt.Printf("      💭 Still need: %dh %dm to complete your weekly goal\n", hours, minutes)
		} else {
			fmt.Printf("      💭 Still need: %dm to complete your weekly goal\n", minutes)
		}
	} else {
		excess := progress.WorkoutHours - progress.Goals.WorkoutGoalHours
		fmt.Printf("      🎉 Goal achieved! You've exceeded by %.1f hours\n", excess)
	}
	
	// Weekly activity summary
	fmt.Printf("\n   📊 This Week Summary:\n")
	fmt.Printf("      🏃 Runs: %d activities\n", progress.RunCount)
	fmt.Printf("      💪 Workouts: %d activities\n", progress.WorkoutCount)
	fmt.Printf("      📈 Total: %d activities\n", progress.TotalActivities)
	
	// Motivational message
	fmt.Printf("\n   💬 %s\n", progress.GetMotivationalMessage())
}

// getProgressDisplay returns a progress bar and status emoji based on percentage
func getProgressDisplay(percent float64) (string, string) {
	var status string
	var bar string
	
	// Determine status emoji
	if percent >= 100 {
		status = "✅ COMPLETED"
	} else if percent >= 75 {
		status = "🟡 ALMOST THERE"
	} else if percent >= 50 {
		status = "🟠 HALFWAY"
	} else if percent >= 25 {
		status = "🔵 GETTING STARTED"
	} else {
		status = "🔴 JUST STARTED"
	}
	
	// Create progress bar (20 characters wide)
	filled := int(percent / 5) // Each character represents 5%
	if filled > 20 {
		filled = 20
	}
	
	bar = "["
	for i := 0; i < 20; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}
	bar += "]"
	
	return status, bar
}
