# Strava Custom Goals 🏃‍♂️

A simple and efficient Go application that fetches and analyzes your Strava activities with custom goal tracking capabilities.

## Features ✨

- 🔐 Secure OAuth token refresh for Strava API access
- 📊 Fetch and display recent activities with comprehensive data
- 🎯 Enhanced activity analysis with calculated fields (pace, distance conversion)
- 🏃‍♂️ Special handling for running activities with pace calculations
- ❤️ Heart rate data display when available
- 📅 Beautiful, emoji-enhanced activity summaries

## Quick Start 🚀

### 1. Setup Strava API
1. Go to [Strava API Settings](https://www.strava.com/settings/api)
2. Create a new application
3. Note down your `Client ID` and `Client Secret`
4. Complete the OAuth flow once to get your `Refresh Token`

### 2. Configure Application
Update the constants in `main.go`:
```go
const (
    CLIENT_ID     = "your_client_id"     
    CLIENT_SECRET = "your_client_secret" 
    REFRESH_TOKEN = "your_refresh_token" 
)
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

## Contributing 🤝

This project welcomes contributions during Hacktoberfest and beyond! Feel free to:
- Add new goal tracking features
- Improve activity analysis
- Enhance the user interface
- Add tests and documentation

## License 📄

Licensed under the MIT License. See [LICENSE](LICENSE) for details.
