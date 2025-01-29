package analyzer

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAnalyzeLogs(t *testing.T) {
	logContent := `[2023-10-01 08:00:00] INFO: Application started.
[2023-10-01 08:00:05] DEBUG: Loading configuration file "config.json".
[2023-10-01 08:00:06] INFO: Configuration loaded successfully.
[2023-10-01 08:00:10] INFO: Database connection established.
[2023-10-01 08:00:15] INFO: User "admin" logged in.
[2023-10-01 08:00:20] DEBUG: Fetching user data for ID 123.
[2023-10-01 08:00:25] INFO: User data retrieved successfully.
[2023-10-01 08:00:30] WARNING: Deprecated API used in module "user_auth".
[2023-10-01 08:00:35] INFO: User "admin" updated profile settings.
[2023-10-01 08:00:40] DEBUG: Sending email notification to user "admin".
[2023-10-01 08:00:45] INFO: Email sent successfully.
[2023-10-01 08:00:50] ERROR: Failed to connect to external service. Retrying in 5 seconds...
[2023-10-01 08:00:55] ERROR: External service connection failed after 3 attempts.
[2023-10-01 08:01:00] CRITICAL: Application shutting down due to critical error.`

	tmpFile, err := os.CreateTemp("", "test_logs_*.go")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err = tmpFile.WriteString(logContent); err != nil {
		t.Fatalf("Failed to write logs to tmp file: %v", err)
	}
	defer tmpFile.Close()

	// Test 1
	t.Run("Total logs count", func(t *testing.T) {
		var stats *LogStats
		stats, err = AnalyzeLogs(tmpFile.Name(), "")
		if err != nil {
			t.Fatalf("Failed to analyze test logs: %v", err)
		}
		expectedCount := 14
		assert.Equal(t, expectedCount, stats.TotalLogs)
	})

	// Test 2
	t.Run("Logs by level", func(t *testing.T) {
		var stats *LogStats
		stats, err = AnalyzeLogs(tmpFile.Name(), "")
		if err != nil {
			t.Fatalf("Failed to analyze test logs: %v", err)
		}
		expectedLogsByLevel := map[string]int{
			"INFO":     7,
			"DEBUG":    3,
			"WARNING":  1,
			"ERROR":    2,
			"CRITICAL": 1,
		}

		for level, expectedCount := range expectedLogsByLevel {
			assert.Equal(t, expectedCount, stats.LogsByLevel[level])
		}
	})

	// Test 3
	t.Run("Filter by log level", func(t *testing.T) {
		var stats *LogStats
		stats, err = AnalyzeLogs(tmpFile.Name(), "ERROR")
		if err != nil {
			t.Fatalf("AnalyzeLogs failed: %v", err)
		}
		expectedLogs := 2
		assert.Equal(t, expectedLogs, stats.LogsByLevel["ERROR"])
	})

	// Test 4
	t.Run("Earliest and latest log", func(t *testing.T) {
		var stats *LogStats
		stats, err = AnalyzeLogs(tmpFile.Name(), "")
		if err != nil {
			t.Fatalf("Failed to analyze test logs: %v", err)
		}

		expectedEarliestLog := time.Date(2023, 10, 1, 8, 0, 0, 0, time.UTC)
		expectedLatestLog := time.Date(2023, 10, 1, 8, 1, 0, 0, time.UTC)
		assert.Equal(t, expectedEarliestLog, stats.EarliestLog)
		assert.Equal(t, expectedLatestLog, stats.LatestLog)
	})

	// Test 5
	t.Run("Error handling", func(t *testing.T) {
		_, err = AnalyzeLogs("notExists", "")
		expectedError := "failed to open log file: open notExists: no such file or directory"
		assert.EqualError(t, err, expectedError)
	})
}
