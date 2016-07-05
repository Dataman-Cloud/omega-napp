package logger

import (
	//"github.com/Dataman-Cloud/omega-napp/config"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"time"
	//"os"
)

func init() {
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{
		ForceColors:      true,
		DisableColors:    false,
		DisableTimestamp: false,
		FullTimestamp:    true,
		DisableSorting:   false,
	})
	// Output to stderr instead of stdout, could also be a file.
	//log.SetOutput(os.Stderr)
	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func Ginrus(logger *log.Logger, timeFormat string, utc bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		if utc {
			end = end.UTC()
		}

		entry := log.WithFields(log.Fields{
			"status":     c.Writer.Status(),
			"method":     c.Request.Method,
			"path":       path,
			"ip":         c.ClientIP(),
			"latency":    latency,
			"user-agent": c.Request.UserAgent(),
			"time":       end.Format(timeFormat),
		})
		_ = entry

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			//entry.Error(c.Errors.String())
			logger.Errorf("%d %s %s %s %s %s %s %s",
				c.Writer.Status(),
				c.Request.Method,
				path,
				c.ClientIP(),
				latency,
				c.Request.UserAgent(),
				end.Format(timeFormat),
				c.Errors.String(),
			)
		} else {
			//entry.Info()
			logger.Infof("%d %s %s %s %s %s %s",
				c.Writer.Status(),
				c.Request.Method,
				path,
				c.ClientIP(),
				latency,
				c.Request.UserAgent(),
				end.Format(timeFormat),
			)
		}
	}
}
