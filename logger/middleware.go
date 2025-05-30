package logger

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

// LoggerMiddleware returns a Gin middleware that logs requests
func LoggerMiddleware(logger *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		// Log request details with query parameters, headers, IP, and duration
		logger.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("ip", c.ClientIP()).            // IP Address
			Int("status", c.Writer.Status()).   // Status Code
			Dur("duration", time.Since(start)). // Duration
			Msg("Handled request")
	}
}
