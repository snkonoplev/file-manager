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
	engine           *gin.Engine
	usersController  *controller.UsersController
	systemController *controller.SystemController
	a                *auth.Auth
	viper            *viper.Viper
	p                *proxy.Proxy
}

func NewRouter(engine *gin.Engine, usersController *controller.UsersController, auth *auth.Auth, viper *viper.Viper, p *proxy.Proxy, systemController *controller.SystemController) *Router {
	return &Router{
		engine:           engine,
		usersController:  usersController,
		systemController: systemController,
		a:                auth,
		viper:            viper,
		p:                p,
	}
}

func (r *Router) MapHandlers() error {

	a, err := r.a.AuthMiddleware()
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

	r.engine.Any("/transmission/*proxyPath", a.MiddlewareFunc(), r.p.Handle)

	api := r.engine.Group("/api")
	{
		api.POST("/login", a.LoginHandler)
		api.GET("/refresh_token", a.RefreshHandler)

		system := api.Group("/system").Use(a.MiddlewareFunc())
		{
			system.GET("/disk-usage", r.systemController.GetDiskUsage)
		}

		users := api.Group("/users").Use(a.MiddlewareFunc())
		{
			users.GET("", r.usersController.GetUsers)
			users.GET(":id", r.usersController.GetUser)
			users.GET("/current", r.usersController.CurrentUser)
			users.POST("", r.usersController.CreteUser)
			users.PUT("", r.usersController.UpdateUser)
			users.DELETE(":id", r.usersController.DeleteUser)
			users.PUT("/change-password", r.usersController.ChangePassword)
		}
	}

	return nil
}
