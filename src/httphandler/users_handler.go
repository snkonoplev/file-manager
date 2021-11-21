package httphandler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snkonoplev/file-manager/mediator"
	"github.com/snkonoplev/file-manager/query"
)

type UsersHandler struct {
	mediator *mediator.Mediator
}

func NewUsersHandler(mediator *mediator.Mediator) *UsersHandler {
	return &UsersHandler{
		mediator: mediator,
	}
}

func (h *UsersHandler) GetUsers(c *gin.Context) {
	result, err := h.mediator.Handle(c.Request.Context(), query.UsersQuery{})
	if err != nil {
		target := &mediator.HandlerError{}
		if errors.As(err, &target) {
			c.String(target.StatusCode, target.Message)
			return
		}

		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, result)
}
