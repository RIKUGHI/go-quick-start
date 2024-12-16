package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type LogFileRotator struct {
	mu       sync.Mutex
	basePath string
	maxAge   time.Duration
	file     *os.File
	curDate  string
}

func NewLogFileRotator(basePath string, maxAge time.Duration) *LogFileRotator {
	logger := &LogFileRotator{
		basePath: basePath,
		maxAge:   maxAge,
	}

	// Ensure the base directory exists
	if err := os.MkdirAll(basePath, 0755); err != nil {
		log.Fatalf("failed to create base directory: %+v", err)
	}

	return logger
}

// Write implements io.Writer and handles log rotation
func (r *LogFileRotator) Write(p []byte) (n int, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Rotate the log if the date has changed
	today := time.Now().Format(time.DateOnly)
	if today != r.curDate {
		if err := r.rotate(); err != nil {
			return 0, fmt.Errorf("log rotation failed: %w", err)
		}
	}

	// Ensure the file is available before writing
	if r.file == nil {
		return 0, fmt.Errorf("log file is not available")
	}

	return r.file.Write(p)
}

// rotate creates a new log file, sets the next rotation time, and cleans up old logs
func (r *LogFileRotator) rotate() error {
	// Close the current file if it exists
	if r.file != nil {
		r.file.Close()
	}

	// Create a new log file
	now := time.Now()
	r.curDate = now.Format(time.DateOnly)
	logFileName := filepath.Join(r.basePath, fmt.Sprintf("%s.log", r.curDate))
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to create log file: %w", err)
	}

	r.file = file

	// Clean up old logs asynchronously
	go r.cleanupOldLogs()

	return nil
}

// cleanupOldLogs removes log files older than maxAge
func (r *LogFileRotator) cleanupOldLogs() {
	files, err := filepath.Glob(filepath.Join(r.basePath, "*.log"))
	if err != nil {
		logrus.Errorf("failed to list log files: %+v", err)
		return
	}

	// Get the cutoff time based on maxAge
	cutoff := time.Now().Add(-r.maxAge)

	// Check each file and delete if older than maxAge
	for _, file := range files {
		if info, err := os.Stat(file); err == nil && info.ModTime().Before(cutoff) {
			if err := os.Remove(file); err != nil {
				logrus.WithError(err).WithField("file", file).Error("Failed to delete old log file")
			} else {
				logrus.WithField("file", file).Info("Deleted old log file")
			}
		}
	}
}

// Close closes the currently open log file
func (r *LogFileRotator) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.file != nil {
		err := r.file.Close()
		r.file = nil
		return err
	}

	return nil
}
