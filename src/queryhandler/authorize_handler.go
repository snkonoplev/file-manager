package queryhandler

import (
	"context"
	"net/http"

	"github.com/snkonoplev/file-manager/db"
	"github.com/snkonoplev/file-manager/mediator"
	"github.com/snkonoplev/file-manager/query"
)

type AuthorizeHandler struct {
	repository db.IUsersRepository
}

func NewAuthorizeHandler(repository db.IUsersRepository) *AuthorizeHandler {
	return &AuthorizeHandler{
		repository: repository,
	}
}

func (h *AuthorizeHandler) Handle(context context.Context, q interface{}) (interface{}, error) {
	if userAuthorizeQuery, ok := q.(query.UserAuthorizeQuery); ok {
		user, err := h.repository.Authorize(context, userAuthorizeQuery.UserName, userAuthorizeQuery.Password)
		if err != nil {
			return nil, &mediator.HandlerError{
				StatusCode: http.StatusInternalServerError,
				Message:    "can't authorize user",
				Err:        err,
			}
		}

		if user.Id == 0 {
			return nil, &mediator.HandlerError{
				StatusCode: http.StatusUnauthorized,
				Message:    "can't find user",
				Err:        err,
			}
		}

		if !user.IsActive {
			return nil, &mediator.HandlerError{
				StatusCode: http.StatusUnauthorized,
				Message:    "user is inactive",
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
