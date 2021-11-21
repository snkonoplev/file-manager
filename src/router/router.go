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

func (r *Router) MapHandlers() {
	users := r.engine.Group("/users")
	{
		users.GET("", r.usersHandler.GetUsers)
	}
}
