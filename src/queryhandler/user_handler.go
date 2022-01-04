package queryhandler

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/snkonoplev/file-manager/db"
	"github.com/snkonoplev/file-manager/mediator"
	"github.com/snkonoplev/file-manager/query"
)

type UserHandler struct {
	repository db.IUsersRepository
}

func NewUserHandler(repository db.IUsersRepository) *UserHandler {
	return &UserHandler{
		repository: repository,
	}
}

func (h *UserHandler) Handle(context context.Context, q interface{}) (interface{}, error) {
	if userQuery, ok := q.(query.UserQuery); ok {
		user, err := h.repository.GetUser(context, userQuery.Id)
		if err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				return nil, &mediator.HandlerError{
					StatusCode: http.StatusNotFound,
					Message:    "can't find user",
				}
			}

			return nil, &mediator.HandlerError{
				StatusCode: http.StatusInternalServerError,
				Message:    "can't get user",
				Err:        err,
			}
		}

		return user, nil
	}

	return nil, &mediator.HandlerError{
		StatusCode: http.StatusInternalServerError,
		Message:    "wrong command type",
	}
}
