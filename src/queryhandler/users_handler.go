package queryhandler

import (
	"context"
	"net/http"

	"github.com/snkonoplev/file-manager/db"
	"github.com/snkonoplev/file-manager/mediator"
)

type UsersHandler struct {
	repository db.IUsersRepository
}

func NewUsersHandler(repository db.IUsersRepository) *UsersHandler {
	return &UsersHandler{
		repository: repository,
	}
}

func (h *UsersHandler) Handle(context context.Context, query interface{}) (interface{}, error) {
	users, err := h.repository.ListUsers(context)

	if err != nil {
		return nil, &mediator.HandlerError{
			StatusCode: http.StatusInternalServerError,
			Message:    "can't get users list",
			Err:        err,
		}
	}

	return users, nil
}
