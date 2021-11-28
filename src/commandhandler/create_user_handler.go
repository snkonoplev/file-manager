package commandhandler

import (
	"context"
	"net/http"
	"time"

	"github.com/snkonoplev/file-manager/command"
	"github.com/snkonoplev/file-manager/db"
	"github.com/snkonoplev/file-manager/entity"
	"github.com/snkonoplev/file-manager/mediator"
	"github.com/snkonoplev/file-manager/security"
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

		if !createUserCommand.IsCallerAdmin {
			return nil, &mediator.HandlerError{
				StatusCode: http.StatusForbidden,
				Message:    "you need admin rights to perform this operation",
			}
		}

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

		passwordHash, err := security.HashPassword(createUserCommand.Password)
		if err != nil {
			return nil, &mediator.HandlerError{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Err:        err,
			}
		}

		user := entity.User{
			Created:  time.Now().UTC().UnixMilli(),
			Name:     createUserCommand.Name,
			Password: passwordHash,
			IsAdmin:  createUserCommand.IsAdmin,
			IsActive: createUserCommand.IsActive,
		}

		userId, err := h.repository.CreateUser(context, user)
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
