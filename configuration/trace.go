package configuration

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("traceId", uuid.New())
	}
}
