package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

var logFile *os.File

// InitLogger initializes the logging system
func InitLogger() {
	var err error

	// Create logs directory if it doesn't exist
	if _, err := os.Stat("./logs"); os.IsNotExist(err) {
		os.MkdirAll("./logs", 0755)
	}

	// Open or create log file
	logFile, err = os.OpenFile("./logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	// Set log output to both console and file
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

// LogInfo logs an info message
func LogInfo(message string) {
	logMessage("INFO", message)
}

// LogError logs an error message
func LogError(message string) {
	logMessage("ERROR", message)
}

// LogWarning logs a warning message
func LogWarning(message string) {
	logMessage("WARNING", message)
}

// logMessage logs a message with a specific level
func logMessage(level, message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("[%s] %s: %s\n", timestamp, level, message)

	// Write to file
	if logFile != nil {
		logFile.WriteString(logEntry)
	}

	// Also print to console
	fmt.Print(logEntry)
}

// CloseLogger closes the log file
func CloseLogger() {
	if logFile != nil {
		logFile.Close()
	}
}