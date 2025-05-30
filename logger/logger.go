package logger

import (
	"fmt"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// SetupLogger initializes the logger with log file configuration
func SetupLogger(logDir string) {
	// รับค่า Pod Name จาก environment variable (ถ้ามี)
	podName := os.Getenv("POD_NAME")
	if podName == "" {
		log.Fatal().Msg("POD_NAME environment variable is not set")
	}

	// สร้าง directory สำหรับ log ถ้ายังไม่มี
	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create log directory")
	}

	// สร้างชื่อไฟล์ log โดยใช้ Pod Name
	logFile := fmt.Sprintf("%s/%s-%s.log", logDir, podName, time.Now().Format("2006-01-02"))

	// Create a new file rotator
	writer, err := rotatelogs.New(
		logFile,
		rotatelogs.WithLinkName(logDir+"/current.log"), // generate symlink to current log file
		rotatelogs.WithMaxAge(7*24*time.Hour),          // files older than 7 days will be deleted
		rotatelogs.WithRotationTime(24*time.Hour),      // rotate every 24 hours
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize log file rotator")
	}

	// สร้าง logger ที่ใช้เฉพาะไฟล์ log
	logger := zerolog.New(writer).With().Timestamp().Logger()
	log.Logger = logger
}

// GetLogger returns the global logger instance
func GetLogger() *zerolog.Logger {
	return &log.Logger
}
