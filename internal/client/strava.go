// Package client provides HTTP client functionality for interacting with the Strava API.
package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"strava-custom-goals/config"
	"strava-custom-goals/internal/models"
)

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
		httpClient:   &http.Client{Timeout: config.RequestTimeout},
	}
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

	resp, err := c.httpClient.Post(config.StravaTokenURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("token request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("token API error %d: %s", resp.StatusCode, string(body))
	}

	var tokenResp models.TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", fmt.Errorf("decode token response: %w", err)
	}

	if tokenResp.AccessToken == "" {
		return "", fmt.Errorf("empty access token received")
	}

	return tokenResp.AccessToken, nil
}

// GetActivities fetches recent activities using the provided access token
func (c *StravaClient) GetActivities(accessToken string) ([]models.Activity, error) {
	url := fmt.Sprintf("%s?per_page=%d&page=1", config.StravaActivitiesURL, config.DefaultPerPage)

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

	var activities []models.Activity
	if err := json.Unmarshal(body, &activities); err != nil {
		return nil, fmt.Errorf("unmarshal activities: %w", err)
	}

	return activities, nil
}
