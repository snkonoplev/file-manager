package router

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/snkonoplev/file-manager/entity"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func GetAuth() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*entity.User); ok {
				return jwt.MapClaims{
					"id": v.Name,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &entity.User{
				Name: claims["id"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
				return &entity.User{
					Name: userID,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*entity.User); ok && v.Name == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.Status(code)
		},
		TokenLookup:   "header: Authorization, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
}
