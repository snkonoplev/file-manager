package db

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/snkonoplev/file-manager/comand"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateUser(context context.Context, user comand.CreateUserCommand) (int64, error) {
	sql := "INSERT INTO users (created, name, password, is_admin) VALUES ($1,$2,$3,$4)"
	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

	result, err := r.db.ExecContext(context, sql, timestamp, user.Name, user.Password, user.IsAdmin)
	if err != nil {
		return 0, fmt.Errorf("can't insert new user %s", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("can't get last inserted id %s", err)
	}

	return id, nil
}

func (r *Repository) CheckUserExists(context context.Context, userName string) (bool, error) {
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
