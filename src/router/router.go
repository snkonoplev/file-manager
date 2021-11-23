package router

import (
	"github.com/gin-gonic/gin"
	"github.com/snkonoplev/file-manager/httphandler"
)

type Router struct {
	engine       *gin.Engine
	usersHandler *httphandler.UsersHandler
}

func NewRouter(engine *gin.Engine, usersHandler *httphandler.UsersHandler) *Router {
	return &Router{
		engine:       engine,
		usersHandler: usersHandler,
	}
}

func (r *Router) MapHandlers() error {

	auth, err := GetAuth()
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
