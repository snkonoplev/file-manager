package db

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/snkonoplev/file-manager/entity"
	"github.com/snkonoplev/file-manager/mediator"
	"github.com/snkonoplev/file-manager/security"
)

//go:generate mockgen -destination ../commandhandler/mock_users_repository_test.go -package commandhandler_test github.com/snkonoplev/file-manager/db IUsersRepository
//go:generate mockgen -destination ../queryhandler/mock_users_repository_test.go -package queryhandler_test github.com/snkonoplev/file-manager/db IUsersRepository
type IUsersRepository interface {
	CreateUser(context context.Context, user entity.User) (int64, error)
	CheckUserExists(context context.Context, userName string) (bool, error)
	ListUsers(context context.Context) ([]entity.User, error)
	Authorize(context context.Context, userName string, password string) (entity.User, error)
	UpdateUser(context context.Context, user entity.User) (entity.User, error)
	DeleteUser(context context.Context, id int64) (int64, error)
	GetUser(context context.Context, id int64) (entity.User, error)
	ChangePassword(context context.Context, name string, password string) (int64, error)
}

type UsersRepository struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) IUsersRepository {
	return &UsersRepository{
		db: db,
	}
}

func (r *UsersRepository) CreateUser(context context.Context, user entity.User) (int64, error) {

	sql := `INSERT INTO users 
			(created, name, password, is_admin, is_active) 
			VALUES (:created,:name,:password,:is_admin,:is_active)`

	result, err := r.db.NamedExecContext(context, sql, user)
	if err != nil {
		return 0, fmt.Errorf("can't insert new user %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("can't get last inserted id %w", err)
	}

	return id, nil
}

func (r *UsersRepository) CheckUserExists(context context.Context, userName string) (bool, error) {
	sql := "SELECT COUNT(*) as count FROM users WHERE name=$1"
	count := 0
	err := r.db.GetContext(context, &count, sql, userName)
	if err != nil {
		return false, fmt.Errorf("checking if user exists failed %w", err)
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func (r *UsersRepository) ListUsers(context context.Context) ([]entity.User, error) {
	sql := "SELECT id, created, last_login, name, is_admin, is_active, password FROM users"
	users := []entity.User{}
	err := r.db.SelectContext(context, &users, sql)
	if err != nil {
		return nil, fmt.Errorf("can't get users list %w", err)
	}

	return users, nil
}

func (r *UsersRepository) Authorize(context context.Context, userName string, password string) (entity.User, error) {
	user := entity.User{}
	sql := "SELECT id, created, last_login, name, is_admin, is_active, password FROM users WHERE name=$1"
	err := r.db.GetContext(context, &user, sql, userName)
	if err != nil {
		return user, err
	}

	if security.CheckPasswordHash(password, user.Password) {
		sql = "UPDATE users SET last_login=:last_login WHERE id=:id"
		_, err := r.db.NamedExecContext(context, sql, map[string]interface{}{
			"last_login": time.Now().UTC().Unix(),
			"id":         user.Id,
		})
		if err != nil {
			return user, err
		}

		return user, nil
	}

	return user, &mediator.HandlerError{
		StatusCode: http.StatusBadRequest,
		Message:    "incorrect password",
	}
}

func (r *UsersRepository) UpdateUser(context context.Context, user entity.User) (entity.User, error) {
	sql := "UPDATE users SET is_active=:is_active, is_admin=:is_admin WHERE id=:id"
	result, err := r.db.NamedExecContext(context, sql, user)
	if err != nil {
		return user, fmt.Errorf("can't update user %w", err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		return user, fmt.Errorf("can't get affected rows %w", err)
	}

	if count == 1 {
		return user, nil
	}

	return user, &mediator.HandlerError{
		StatusCode: http.StatusBadRequest,
		Message:    "can't find user",
	}
}

func (r *UsersRepository) DeleteUser(context context.Context, id int64) (int64, error) {
	sql := "DELETE FROM users WHERE id=?"
	result, err := r.db.ExecContext(context, sql, id)
	if err != nil {
		return id, fmt.Errorf("can't delete user %w", err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		return id, fmt.Errorf("can't get affected rows %w", err)
	}

	if count == 1 {
		return id, nil
	}

	return id, &mediator.HandlerError{
		StatusCode: http.StatusBadRequest,
		Message:    "can't find user",
	}
}

func (r *UsersRepository) GetUser(context context.Context, id int64) (entity.User, error) {
	user := entity.User{}
	sql := "SELECT * FROM users WHERE id=?"
	err := r.db.Get(&user, sql, id)
	if err != nil {
		return user, fmt.Errorf("can't get user %w", err)
	}

	return user, nil
}

func (r *UsersRepository) ChangePassword(context context.Context, name string, password string) (int64, error) {
	sql := "UPDATE users SET password=? WHERE name=?"
	result, err := r.db.ExecContext(context, sql, password, name)
	if err != nil {
		return 0, fmt.Errorf("can't change password %w", err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("can't get affected rows %w", err)
	}

	if count == 1 {
		return count, nil
	}

	return count, &mediator.HandlerError{
		StatusCode: http.StatusBadRequest,
		Message:    "can't find user",
	}
}
