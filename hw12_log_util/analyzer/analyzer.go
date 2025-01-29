package analyzer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type LogStats struct {
	TotalLogs    int            // Общее кол-во логов
	LogsByLevel  map[string]int // Кол-во логов по уровню
	EarliestLog  time.Time      // Самая ранняя запись в логах
	LatestLog    time.Time      // Самая поздняя запись в логах
	ErrorLogs    []string       // Список сообщений с уровнем ERROR
	CriticalLogs []string       // Список сообщений с уровнем CRITICAL
}

func AnalyzeLogs(filepath, logLevel string) (*LogStats, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}
	defer file.Close()

	stats := &LogStats{
		LogsByLevel: make(map[string]int),
	}

	scanner := bufio.NewScanner(file)
	timeFormat := "2006-01-02 15:04:05"

	for scanner.Scan() {
		line := scanner.Text()
		stats.TotalLogs++

		parts := strings.SplitN(line, "]", 2)
		if len(parts) != 2 {
			continue
		}
		timeStampStr := strings.TrimPrefix(parts[0], "[")
		timestamp, err := time.Parse(timeFormat, timeStampStr)
		if err != nil {
			continue
		}

		if stats.EarliestLog.IsZero() || timestamp.Before(stats.EarliestLog) {
			stats.EarliestLog = timestamp
		}
		if stats.LatestLog.IsZero() || timestamp.After(stats.LatestLog) {
			stats.LatestLog = timestamp
		}

		logParts := strings.SplitN(strings.TrimSpace(parts[1]), ":", 2)
		if len(logParts) != 2 {
			continue
		}

		level := strings.ToUpper(strings.TrimSpace(logParts[0]))
		message := strings.TrimSpace(logParts[1])

		if logLevel != "" && level != logLevel {
			continue
		}

		stats.LogsByLevel[level]++

		if level == "ERROR" {
			stats.ErrorLogs = append(stats.ErrorLogs, message)
		}
		if level == "CRITICAL" {
			stats.CriticalLogs = append(stats.CriticalLogs, message)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading log file: %w", err)
	}

	return stats, nil
}
