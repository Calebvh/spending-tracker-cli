package logger

import (
	"fmt"
	"strconv"
	"time"
)

const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorBlue   = "\033[34m"
	colorYellow = "\033[33m"
	formatBold  = "\033[1m"
	formatReset = "\033[0m"
)

// LogLevel represents the log level.
type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarningLevel
	ErrorLevel
)

var (
	logLevel    LogLevel
	debugMode   bool
	warningMode bool
)

// SetLogLevel sets the log level.
func SetLogConfig(LogLevel string, DebugMode bool, WarningMode bool) {

	switch LogLevel {
	case "DEBUG":
		logLevel = DebugLevel
	case "INFO":
		logLevel = InfoLevel
	case "WARNING":
		logLevel = WarningLevel
	case "ERROR":
		logLevel = ErrorLevel
	default:
		logLevel = InfoLevel
	}
	debugMode = DebugMode
	warningMode = WarningMode

	Error("Log level set to " + LogLevel + ". Debug mode set to " + strconv.FormatBool(DebugMode) + ". Warning mode set to " + strconv.FormatBool(WarningMode))
}

func log(level string, message string, color string) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s%s%s [%s] %s%s\n", formatBold, color, currentTime, level, formatReset, message)
}

// debug prints a debug message.
func Debug(message string) {
	if logLevel <= DebugLevel && debugMode {
		log("DEBUG", message, colorBlue)
	}
}

// info prints an info message.
func Info(message string) {
	if logLevel <= InfoLevel {
		log("INFO", message, colorGreen)
	}
}

// warning prints a warning message.
func Warning(message string) {
	if logLevel <= WarningLevel && warningMode {
		log("WARNING", message, colorYellow)
	}
}

// err prints an error message.
func Error(message string) {
	if logLevel <= ErrorLevel {
		log("ERROR", message, colorRed)
	}
}
