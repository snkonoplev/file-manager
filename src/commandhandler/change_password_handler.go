package commandhandler

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/snkonoplev/file-manager/command"
	"github.com/snkonoplev/file-manager/db"
	"github.com/snkonoplev/file-manager/mediator"
	"github.com/snkonoplev/file-manager/security"
)

type ChangePasswordHandler struct {
	repository *db.UsersRepository
}

func NewChangePasswordHandler(repository *db.UsersRepository) *ChangePasswordHandler {
	return &ChangePasswordHandler{
		repository: repository,
	}
}

func (h *ChangePasswordHandler) Handle(context context.Context, c interface{}) (interface{}, error) {
	if changePasswordCommand, ok := c.(command.ChangePasswordCommand); ok {

		if !changePasswordCommand.IsCallerAdmin {
			return nil, &mediator.HandlerError{
				StatusCode: http.StatusForbidden,
				Message:    "you need admin rights to perform this operation",
			}
		}

		_, err := h.repository.Authorize(context, changePasswordCommand.Name, changePasswordCommand.PreviousPassword)
		if err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				return nil, &mediator.HandlerError{
					StatusCode: http.StatusNotFound,
					Message:    "can't find user",
				}
			}

			target := &mediator.HandlerError{}
			if errors.As(err, &target) {
				return nil, target
			}

			return nil, &mediator.HandlerError{
				StatusCode: http.StatusInternalServerError,
				Message:    "can't check previous password",
				Err:        err,
			}
		}

		passwordHash, err := security.HashPassword(changePasswordCommand.NewPassword)
		if err != nil {
			return nil, &mediator.HandlerError{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Err:        err,
			}
		}

		id, err := h.repository.ChangePassword(context, changePasswordCommand.Name, passwordHash)
		if err != nil {

			target := &mediator.HandlerError{}
			if errors.As(err, &target) {
				return nil, target
			}

			return nil, &mediator.HandlerError{
				StatusCode: http.StatusInternalServerError,
				Message:    "can't change password",
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
