package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func NewDbConnection(filepath string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", filepath)
	if err != nil {
		return nil, fmt.Errorf("can't create db connection %s", err)
	}

	return db, nil
}
