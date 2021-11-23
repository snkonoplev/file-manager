package auth

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/snkonoplev/file-manager/entity"
	"github.com/snkonoplev/file-manager/mediator"
	"github.com/snkonoplev/file-manager/query"
	"github.com/spf13/viper"
)

type Claim struct {
	UserName string
}

type Auth struct {
	secret string
	m      *mediator.Mediator
}

func NewAuth(viper *viper.Viper, mediator *mediator.Mediator) *Auth {
	return &Auth{
		secret: viper.GetString("JWT_SECRET"),
		m:      mediator,
	}
}

var identityKey string = "user_name"

func (a *Auth) AuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "file-manager",
		Key:         []byte(a.secret),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(entity.UserFull); ok {
				return jwt.MapClaims{
					identityKey: v.Name,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &Claim{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var query query.UserAuthorizeQuery
			if err := c.ShouldBind(&query); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			user, err := a.m.Handle(c.Request.Context(), query)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			return user, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*Claim); ok && v.UserName == "admin" {
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
