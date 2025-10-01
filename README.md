# Strava Custom Goals 🏃‍♂️

A simple and efficient Go application that fetches and analyzes your Strava activities with custom goal tracking capabilities.

## Features ✨

- 🔐 Secure OAuth token refresh for Strava API access
- 📊 Fetch and display recent activities with comprehensive data
- 🎯 **Weekly Goals Tracking** - Set and track running and workout goals
- 🏃‍♂️ Running goal tracking with distance progress (km per week)
- 💪 Workout goal tracking with time progress (hours per week)
- 📈 Progress visualization with percentages and motivational messages
- 🎯 Enhanced activity analysis with calculated fields (pace, distance conversion)
- ❤️ Heart rate data display when available
- 📅 Beautiful, emoji-enhanced activity summaries

## Quick Start 🚀

### 1. Setup Strava API
1. Go to [Strava API Settings](https://www.strava.com/settings/api)
2. Create a new application
3. Note down your `Client ID` and `Client Secret`
4. Complete the OAuth flow once to get your `Refresh Token`

### 2. Configure Application
Copy the example environment file and update it with your credentials:
```bash
cp .env.example .env
```

Then edit `.env` with your actual Strava API credentials and weekly goals:
```env
# Strava API Configuration
STRAVA_CLIENT_ID=your_actual_client_id
STRAVA_CLIENT_SECRET=your_actual_client_secret
STRAVA_REFRESH_TOKEN=your_actual_refresh_token

# Weekly Goals Configuration
WEEKLY_RUNNING_GOAL_KM=10      # Target: 10km of running per week
WEEKLY_WORKOUT_GOAL_HOURS=3    # Target: 3 hours of workouts per week
```

### 3. Run the Application
```bash
go run main.go
```

## Sample Output 📈

```
🚀 Strava Custom Goals Tracker Starting...
📡 Authenticating with Strava API...
✅ Successfully authenticated
📊 Fetching recent activities...
✅ Retrieved 5 activities
🎯 Calculating weekly goals progress...

🎯 === WEEKLY GOALS PROGRESS ===
   🏃‍♂️ Running Target: 8.5 km / 10.0 km (85.0%)
      [█████████████████░░░] � ALMOST THERE
      �💭 Still need: 1.5 km to complete your weekly goal

   💪 Workout Target: 2.5 hours / 3.0 hours (83.3%)
      [████████████████░░░░] 🟡 ALMOST THERE
      💭 Still need: 30m to complete your weekly goal

   📊 This Week Summary:
      🏃 Runs: 3 activities
      💪 Workouts: 2 activities
      📈 Total: 5 activities

   💬 🔥 You're over halfway to both goals! Keep pushing!

🏃‍♂️ === RECENT ACTIVITIES ===

📈 Activity 1
   🏷️  Name: Morning Run
   🎯 Type: Run
   📏 Distance: 5.23 km
   ⏱️  Moving Time: 25m 14s
   🏃 Average Pace: 4:49 min/km
   ❤️  Avg Heart Rate: 165 bpm
   👍 Kudos: 3
   📅 Date: Oct 1, 2025 07:30
```
🚀 Strava Custom Goals Tracker Starting...
📡 Authenticating with Strava API...
✅ Successfully authenticated
📊 Fetching recent activities...
✅ Retrieved 5 activities

🏃‍♂️ === RECENT ACTIVITIES ===

📈 Activity 1
   🏷️  Name: Morning Run
   🎯 Type: Run
   📏 Distance: 5.23 km
   ⏱️  Moving Time: 25m 14s
   🏃 Average Pace: 4:49 min/km
   ❤️  Avg Heart Rate: 165 bpm
   👍 Kudos: 3
   📅 Date: Oct 1, 2025 07:30
```

## Security 🔒

- **Never commit your `.env` file**: The `.env` file contains sensitive API credentials and is automatically ignored by git
- **Use `.env.example`**: This file shows the required environment variables without exposing actual values
- **Rotate credentials**: If you accidentally expose your credentials, regenerate them in your Strava API settings

## Contributing 🤝

This project welcomes contributions during Hacktoberfest and beyond! Feel free to:
- Add new goal tracking features
- Improve activity analysis
- Enhance the user interface
- Add tests and documentation

## License 📄

Licensed under the MIT License. See [LICENSE](LICENSE) for details.
