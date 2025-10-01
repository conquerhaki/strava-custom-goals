// Package logger provides structured logging utilities for the application
package logger

import (
	"fmt"
	"log"
	"os"
)

// Logger levels
const (
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
	DEBUG = "DEBUG"
)

// Logger wraps the standard logger with structured logging
type Logger struct {
	*log.Logger
	level string
}

// NewLogger creates a new logger instance
func NewLogger(level string) *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "", log.LstdFlags),
		level:  level,
	}
}

// Info logs an info message
func (l *Logger) Info(msg string, args ...interface{}) {
	l.logMessage(INFO, msg, args...)
}

// Warn logs a warning message
func (l *Logger) Warn(msg string, args ...interface{}) {
	l.logMessage(WARN, msg, args...)
}

// Error logs an error message
func (l *Logger) Error(msg string, args ...interface{}) {
	l.logMessage(ERROR, msg, args...)
}

// Debug logs a debug message
func (l *Logger) Debug(msg string, args ...interface{}) {
	if l.level == DEBUG {
		l.logMessage(DEBUG, msg, args...)
	}
}

// logMessage formats and logs a message
func (l *Logger) logMessage(level, msg string, args ...interface{}) {
	emoji := getEmoji(level)
	formatted := fmt.Sprintf(msg, args...)
	l.Printf("%s [%s] %s", emoji, level, formatted)
}

// getEmoji returns an emoji for the log level
func getEmoji(level string) string {
	switch level {
	case INFO:
		return "ℹ️"
	case WARN:
		return "⚠️"
	case ERROR:
		return "❌"
	case DEBUG:
		return "🐛"
	default:
		return "📝"
	}
}
