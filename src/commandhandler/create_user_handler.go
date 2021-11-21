package commandhandler

import (
	"context"
	"net/http"

	"github.com/snkonoplev/file-manager/command"
	"github.com/snkonoplev/file-manager/db"
	"github.com/snkonoplev/file-manager/mediator"
)

type CreateUserHandler struct {
	repository *db.UsersRepository
}

func NewCreateUserHandler(repository *db.UsersRepository) *CreateUserHandler {
	return &CreateUserHandler{
		repository: repository,
	}
}

func (h *CreateUserHandler) Handle(context context.Context, c interface{}) (interface{}, error) {
	if createUserCommand, ok := c.(command.CreateUserCommand); ok {
		exists, err := h.repository.CheckUserExists(context, createUserCommand.Name)
		if err != nil {
			return nil, &mediator.HandlerError{
				StatusCode: http.StatusInternalServerError,
				Message:    "can't check user if exists",
				Err:        err,
			}
		}

		if exists {
			return nil, &mediator.HandlerError{
				StatusCode: http.StatusBadRequest,
				Message:    "user already exists",
				Err:        err,
			}
		}

		userId, err := h.repository.CreateUser(context, createUserCommand)
		if err != nil {
			return nil, &mediator.HandlerError{
				StatusCode: http.StatusInternalServerError,
				Message:    "can't create new user",
				Err:        err,
			}
		}

		return userId, nil
	}

	return nil, &mediator.HandlerError{
		StatusCode: http.StatusInternalServerError,
		Message:    "wrong command type",
	}
}
