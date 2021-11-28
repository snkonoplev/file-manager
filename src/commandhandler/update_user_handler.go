package commandhandler

import (
	"context"
	"errors"
	"net/http"

	"github.com/snkonoplev/file-manager/command"
	"github.com/snkonoplev/file-manager/db"
	"github.com/snkonoplev/file-manager/entity"
	"github.com/snkonoplev/file-manager/mediator"
)

type UpdateUserHandler struct {
	repository *db.UsersRepository
}

func NewUpdateUserHandler(repository *db.UsersRepository) *UpdateUserHandler {
	return &UpdateUserHandler{
		repository: repository,
	}
}

func (h *UpdateUserHandler) Handle(context context.Context, c interface{}) (interface{}, error) {
	if updateUserCommand, ok := c.(command.UpdateUserCommand); ok {

		if !updateUserCommand.IsCallerAdmin {
			return nil, &mediator.HandlerError{
				StatusCode: http.StatusForbidden,
				Message:    "you need admin rights to perform this operation",
			}
		}

		user := entity.User{
			Id:       updateUserCommand.Id,
			IsAdmin:  updateUserCommand.IsAdmin,
			IsActive: updateUserCommand.IsActive,
		}

		user, err := h.repository.UpdateUser(context, user)

		if err != nil {

			target := &mediator.HandlerError{}
			if errors.As(err, &target) {
				return nil, target
			}

			return nil, &mediator.HandlerError{
				StatusCode: http.StatusInternalServerError,
				Message:    "can't update user",
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
