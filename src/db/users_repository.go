package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

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

func (r *UsersRepository) CreateUser(context context.Context, user entity.User) (int64, error) {

	sql := "INSERT INTO users (created, name, password, is_admin) VALUES (:created,:name,:password,:is_admin)"

	result, err := r.db.NamedExecContext(context, sql, user)
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
	sql := "SELECT id, created, last_login, name, is_admin, password FROM users"
	users := []entity.User{}
	err := r.db.SelectContext(context, &users, sql)
	if err != nil {
		return nil, fmt.Errorf("can't get users list %s", err)
	}

	return users, nil
}

func (r *UsersRepository) Authorize(context context.Context, userName string, password string) (entity.User, error) {
	user := entity.User{}
	sql := "SELECT id, created, last_login, name, is_admin, password FROM users WHERE name=$1"
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

	return user, fmt.Errorf("incorrect password")
}

func (r *UsersRepository) UpdateUser(context context.Context, user entity.User) (entity.User, error) {
	sql := "UPDATE users SET name=:name, password=:password, is_admin=:is_admin WHERE id=:id"
	result, err := r.db.NamedExecContext(context, sql, user)
	if err != nil {
		return user, fmt.Errorf("can't update user %s", err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		return user, fmt.Errorf("can't get affected rows %s", err)
	}

	if count == 1 {
		return user, nil
	}

	return user, fmt.Errorf("can't find user with id %d", user.Id)
}

func (r *UsersRepository) DeleteUser(context context.Context, id int64) (int64, error) {
	sql := "DELETE FROM users WHERE id=:id"
	result, err := r.db.NamedExecContext(context, sql, id)
	if err != nil {
		return id, fmt.Errorf("can't delete user %s", err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		return id, fmt.Errorf("can't get affected rows %s", err)
	}

	if count == 1 {
		return id, nil
	}

	return id, fmt.Errorf("can't find user with id %d", id)
}
