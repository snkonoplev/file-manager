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
	UserId  int64
	IsAdmin bool
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

var Claims string = "claims"
var IdentityKey string = "user_id"
var IsAdminKey string = "is_admin"

// @Id Login
// @Summary Get access token
// @Accept  json
// @Produce  json
// @Param Body body query.UserAuthorizeQuery true "User"
// @Router /api/login [post]
// @Success 200 {object} map[string]string
// @Failure 401 {string} string
// @Tags Auth
func (a *Auth) AuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "file-manager",
		Key:         []byte(a.secret),
		Timeout:     time.Hour,
		IdentityKey: Claims,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(entity.User); ok {
				return jwt.MapClaims{
					IdentityKey: v.Id,
					IsAdminKey:  v.IsAdmin,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			userClaims := Claim{
				UserId:  int64(claims[IdentityKey].(float64)),
				IsAdmin: claims[IsAdminKey].(bool),
			}

			c.Set(Claims, userClaims)

			return &userClaims
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var query query.UserAuthorizeQuery
			if err := c.ShouldBind(&query); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}

			user, err := a.m.Handle(c, query)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			return user, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*Claim); ok {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.String(code, "Unauthorized")
		},
		TokenLookup:   "header: Authorization, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
}
