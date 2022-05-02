package configuration

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := logrus.StandardLogger()

		t := time.Now()
		traceId, _ := c.Get("traceId")

		logger.WithContext(c.Request.Context()).WithFields(logrus.Fields{
			"ip":        c.ClientIP(),
			"method":    c.Request.Method,
			"path":      c.Request.URL.Path,
			"proto":     c.Request.Proto,
			"userAgent": c.Request.UserAgent(),
			"traceId":   traceId,
		}).Info("Start executing...")

		c.Next()

		logger.WithContext(c.Request.Context()).WithFields(logrus.Fields{
			"traceId": traceId,
			"latency": time.Since(t),
			"status":  c.Writer.Status(),
		}).Info("Stop executing...")
	}
}
