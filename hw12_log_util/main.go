package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/Krovaldo/OtusHW/hw12_log_util/analyzer"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Failed to load .env file: %v", err)
		return
	}

	filePath := flag.String("file", "", "Path to the log file")
	logLevel := flag.String("level", "", "Log level to analyze (optional)")
	outputPath := flag.String("output", "", "Path to the output file (optional)")
	flag.Parse()

	if *filePath == "" {
		*filePath = os.Getenv("LOG_ANALYZER_FILE")
	}

	if *filePath == "" {
		log.Fatal("You must provide a file path")
	}

	stats, err := analyzer.AnalyzeLogs(*filePath, *logLevel)
	if err != nil {
		log.Fatalf("Error analyzing logs: %v", err)
	}

	output := os.Stdout
	if *outputPath != "" {
		var file *os.File
		file, err = os.Create(*outputPath)
		if err != nil {
			log.Fatalf("Error creating output file: %v", err)
		}
		defer file.Close()
		output = file
	}

	fmt.Fprintf(output, "Log statistics:\n")
	fmt.Fprintf(output, "Total logs: %v\n", stats.TotalLogs)
	fmt.Fprintf(output, "Earliest log: %s\n", stats.EarliestLog.Format("2006-01-02 15:04:05"))
	fmt.Fprintf(output, "Latest log: %s\n", stats.LatestLog.Format("2006-01-02 15:04:05"))
	fmt.Fprintf(output, "Logs by level:\n")
	for level, count := range stats.LogsByLevel {
		fmt.Fprintf(output, "    %s: %v\n", level, count)
	}
	if len(stats.ErrorLogs) > 0 {
		fmt.Fprintf(output, "Error logs:\n")
		for _, msg := range stats.ErrorLogs {
			fmt.Fprintf(output, "    - %s\n", msg)
		}
	}
	if len(stats.CriticalLogs) > 0 {
		fmt.Fprintf(output, "Critical logs:\n")
		for _, msg := range stats.CriticalLogs {
			fmt.Fprintf(output, "    - %s\n", msg)
		}
	}
}
