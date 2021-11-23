package router

import (
	"github.com/gin-gonic/gin"
	"github.com/snkonoplev/file-manager/auth"
	"github.com/snkonoplev/file-manager/httphandler"
)

type Router struct {
	engine       *gin.Engine
	usersHandler *httphandler.UsersHandler
	a            *auth.Auth
}

func NewRouter(engine *gin.Engine, usersHandler *httphandler.UsersHandler, auth *auth.Auth) *Router {
	return &Router{
		engine:       engine,
		usersHandler: usersHandler,
		a:            auth,
	}
}

func (r *Router) MapHandlers() error {

	auth, err := r.a.AuthMiddleware()
	if err != nil {
		return err
	}

	users := r.engine.Group("/users").Use(auth.MiddlewareFunc())
	{
		users.GET("", r.usersHandler.GetUsers)
	}

	r.engine.POST("/login", auth.LoginHandler)
	r.engine.GET("/refresh_token", auth.RefreshHandler)

	return nil
}
