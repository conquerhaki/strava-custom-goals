// Package cache provides simple file-based caching for Strava data
package cache

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"strava-custom-goals/internal/models"
)

// Cache handles local data storage
type Cache struct {
	cacheDir string
}

// NewCache creates a new cache instance
func NewCache() *Cache {
	homeDir, _ := os.UserHomeDir()
	cacheDir := filepath.Join(homeDir, ".strava-goals-cache")
	os.MkdirAll(cacheDir, 0755)

	return &Cache{
		cacheDir: cacheDir,
	}
}

// CacheData represents cached activity data
type CacheData struct {
	Activities []models.Activity `json:"activities"`
	Timestamp  time.Time         `json:"timestamp"`
}

// SaveActivities saves activities to cache
func (c *Cache) SaveActivities(activities []models.Activity) error {
	data := CacheData{
		Activities: activities,
		Timestamp:  time.Now(),
	}

	file := filepath.Join(c.cacheDir, "activities.json")
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal cache data: %w", err)
	}

	return os.WriteFile(file, jsonData, 0644)
}

// LoadActivities loads activities from cache if not expired
func (c *Cache) LoadActivities(maxAge time.Duration) ([]models.Activity, bool) {
	file := filepath.Join(c.cacheDir, "activities.json")

	data, err := os.ReadFile(file)
	if err != nil {
		return nil, false // Cache miss
	}

	var cacheData CacheData
	if err := json.Unmarshal(data, &cacheData); err != nil {
		return nil, false // Invalid cache
	}

	// Check if cache is expired
	if time.Since(cacheData.Timestamp) > maxAge {
		return nil, false // Cache expired
	}

	return cacheData.Activities, true
}

// ClearCache removes all cached data
func (c *Cache) ClearCache() error {
	return os.RemoveAll(c.cacheDir)
}
