package db

import (
	"context"
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
	sql := "INSERT INTO users (created, name, password) VALUES ($1,$2,$3)"
	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

	result, err := r.db.ExecContext(context, sql, timestamp, user.Name, user.Password)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
