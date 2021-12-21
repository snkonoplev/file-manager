package commandhandler_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/snkonoplev/file-manager/command"
	"github.com/snkonoplev/file-manager/commandhandler"
	"github.com/snkonoplev/file-manager/entity"
	"github.com/snkonoplev/file-manager/security"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := NewMockIUsersRepository(ctrl)
	var createdUser entity.User

	com := command.CreateUserCommand{
		Name:          "Test",
		Password:      "123",
		IsAdmin:       true,
		IsActive:      true,
		IsCallerAdmin: true,
	}

	call := mock.EXPECT().
		CheckUserExists(gomock.Any(), gomock.Eq(com.Name)).
		Return(false, nil)

	mock.EXPECT().
		CreateUser(gomock.Any(), gomock.Any()).
		DoAndReturn(func(context context.Context, user entity.User) (int64, error) {
			createdUser = user
			return int64(1), nil
		}).
		After(call)

	h := commandhandler.NewCreateUserHandler(mock)
	_, err := h.Handle(context.Background(), com)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, createdUser.Name, com.Name, "Name are not equal")
	assert.Equal(t, createdUser.IsActive, com.IsActive, "IsActive are not equal")
	assert.Equal(t, createdUser.IsAdmin, com.IsAdmin, "IsAdmin are not equal")
	assert.True(t, security.CheckPasswordHash(com.Password, createdUser.Password), "Wrong password hash")
}
