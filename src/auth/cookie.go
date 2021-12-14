package auth

import (
	"github.com/gin-gonic/gin"
)

func SetCookie(c *gin.Context) {
	if header, ok := c.Request.Header["Authorization"]; ok {
		if len(header) > 0 {
			c.SetCookie("jwt", header[0][7:], 3600, "/", c.Request.URL.Hostname(), false, false)
		}
	}
}
