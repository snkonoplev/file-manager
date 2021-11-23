package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/snkonoplev/file-manager/command"
	"github.com/snkonoplev/file-manager/entity"
	"github.com/snkonoplev/file-manager/security"
)

type UsersRepository struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) *UsersRepository {
	return &UsersRepository{
		db: db,
	}
}

func (r *UsersRepository) CreateUser(context context.Context, user command.CreateUserCommand) (int64, error) {
	passwordHash, err := security.HashPassword(user.Password)
	if err != nil {
		return 0, fmt.Errorf("can't calculate password hash %s", err)
	}

	sql := "INSERT INTO users (created, name, password, is_admin) VALUES ($1,$2,$3,$4)"
	result, err := r.db.ExecContext(context, sql, time.Now().UTC().Unix(), user.Name, passwordHash, user.IsAdmin)
	if err != nil {
		return 0, fmt.Errorf("can't insert new user %s", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("can't get last inserted id %s", err)
	}

	return id, nil
}

func (r *UsersRepository) CheckUserExists(context context.Context, userName string) (bool, error) {
	sql := "SELECT COUNT(*) as count FROM users WHERE name=$1"
	count := 0
	err := r.db.GetContext(context, &count, sql, userName)
	if err != nil {
		return false, fmt.Errorf("checking if user exists failed %s", err)
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func (r *UsersRepository) ListUsers(context context.Context) ([]entity.User, error) {
	sql := "SELECT id, created, last_login, name, is_admin FROM users"
	users := []entity.User{}
	err := r.db.SelectContext(context, &users, sql)
	if err != nil {
		return nil, fmt.Errorf("can't get users list %s", err)
	}

	return users, nil
}

func (r *UsersRepository) Authorize(context context.Context, userName string, password string) (entity.UserFull, error) {
	user := entity.UserFull{}
	sql := "SELECT id, created, last_login, name, is_admin, password FROM users WHERE name=$1"
	err := r.db.GetContext(context, &user, sql, userName)
	if err != nil {
		return user, err
	}

	if security.CheckPasswordHash(password, user.Password) {
		return user, nil
	}

	return user, fmt.Errorf("incorrect password")
}
