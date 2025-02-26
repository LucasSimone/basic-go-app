package main

import (
	"log"
	"log/slog"
	"os"
	"strconv"
)

// A custom logger that will log outputs to the console if the enviroment variable DEBUG
// is True. Otherwise it will log a JSON output to a log file
func NewCustomLogger(filePath string) *slog.Logger {

	// Get the Debug state
	isDebug, err := strconv.ParseBool(getEnv("DEBUG", "False"))
	if err != nil {
		log.Fatalf("failed to check debug mode: %v", err)
	}

	// Text logger that is used to output logs to the console when in debug mode
	textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: false,
	})
	textLogger := slog.New(textHandler)

	if isDebug {
		return textLogger
	}

	// Debug is false log to log file
	logFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}

	// JSON handler that is used to output logs to log files when running in production
	jsonHandler := slog.NewJSONHandler(logFile, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	})
	jsonLogger := slog.New(jsonHandler)

	return jsonLogger

}
