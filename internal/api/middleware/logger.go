package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"time"
)

func Logger(c *gin.Context) {
	start := time.Now()

	c.Next()

	var record *zerolog.Event
	if c.Errors.Last() == nil {
		record = log.Info()
	} else {
		record = log.Error().Err(c.Errors.Last())
	}

	record.
		Str("method", c.Request.Method).
		Str("path", c.Request.URL.Path).
		Str("query", c.Request.URL.RawQuery).
		Str("user-agent", c.Request.UserAgent()).
		Int("status", c.Writer.Status())

	// Capture latency
	record.Dur("latency", time.Since(start))

	// Send log entry
	record.Send()
}
