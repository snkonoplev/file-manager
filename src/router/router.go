package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/snkonoplev/file-manager/auth"
	"github.com/snkonoplev/file-manager/controller"
	"github.com/snkonoplev/file-manager/proxy"
	"github.com/spf13/viper"
)

type Router struct {
	engine          *gin.Engine
	usersController *controller.UsersController
	a               *auth.Auth
	viper           *viper.Viper
}

func NewRouter(engine *gin.Engine, usersController *controller.UsersController, auth *auth.Auth, viper *viper.Viper) *Router {
	return &Router{
		engine:          engine,
		usersController: usersController,
		a:               auth,
		viper:           viper,
	}
}

func (r *Router) MapHandlers() error {

	auth, err := r.a.AuthMiddleware()
	if err != nil {
		return err
	}

	if r.viper.GetString("GIN_MODE") == "debug" {
		config := cors.DefaultConfig()
		config.AllowOrigins = []string{"http://localhost:8080"}
		config.AddAllowHeaders("Authorization")
		config.AllowCredentials = true

		r.engine.Use(cors.New(config))
	}

	r.engine.Any("/transmission/*proxyPath", proxy.Proxy)

	r.engine.POST("/login", auth.LoginHandler)
	r.engine.GET("/refresh_token", auth.RefreshHandler)

	users := r.engine.Group("/users").Use(auth.MiddlewareFunc())
	{
		users.GET("", r.usersController.GetUsers)
		users.GET(":id", r.usersController.GetUser)
		users.GET("/current", r.usersController.CurrentUser)
		users.POST("", r.usersController.CreteUser)
		users.PUT("", r.usersController.UpdateUser)
		users.DELETE(":id", r.usersController.DeleteUser)
	}

	return nil
}
