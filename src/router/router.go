package router

import (
	"github.com/gin-gonic/gin"
	"github.com/snkonoplev/file-manager/auth"
	"github.com/snkonoplev/file-manager/controller"
)

type Router struct {
	engine          *gin.Engine
	usersController *controller.UsersController
	a               *auth.Auth
}

func NewRouter(engine *gin.Engine, usersController *controller.UsersController, auth *auth.Auth) *Router {
	return &Router{
		engine:          engine,
		usersController: usersController,
		a:               auth,
	}
}

func (r *Router) MapHandlers() error {

	auth, err := r.a.AuthMiddleware()
	if err != nil {
		return err
	}

	r.engine.POST("/login", auth.LoginHandler)
	r.engine.GET("/refresh_token", auth.RefreshHandler)

	users := r.engine.Group("/users").Use(auth.MiddlewareFunc())
	{
		users.GET("", r.usersController.GetUsers)
		users.POST("", r.usersController.CreteUser)
		users.PUT("", r.usersController.UpdateUser)
	}

	return nil
}
