package commandhandler

import (
	"context"
	"errors"
	"net/http"

	"github.com/snkonoplev/file-manager/command"
	"github.com/snkonoplev/file-manager/db"
	"github.com/snkonoplev/file-manager/mediator"
)

type DeleteUserHandler struct {
	repository *db.UsersRepository
}

func NewDeleteUserHandler(repository *db.UsersRepository) *DeleteUserHandler {
	return &DeleteUserHandler{
		repository: repository,
	}
}

func (h *DeleteUserHandler) Handle(context context.Context, c interface{}) (interface{}, error) {
	if deleteUserCommand, ok := c.(command.DeleteUserCommand); ok {

		if !deleteUserCommand.IsCallerAdmin {
			return nil, &mediator.HandlerError{
				StatusCode: http.StatusForbidden,
				Message:    "you need admin rights to perform this operation",
			}
		}

		id, err := h.repository.DeleteUser(context, deleteUserCommand.Id)
		if err != nil {

			target := &mediator.HandlerError{}
			if errors.As(err, &target) {
				return nil, target
			}

			return nil, &mediator.HandlerError{
				StatusCode: http.StatusInternalServerError,
				Message:    "can't delete user",
				Err:        err,
			}
		}

		return id, nil
	}
	return nil, &mediator.HandlerError{
		StatusCode: http.StatusInternalServerError,
		Message:    "wrong command type",
	}
}
