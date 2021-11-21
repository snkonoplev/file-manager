package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func NewDbConnection() *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", "manager.db")
	if err != nil {
		panic(err)
	}

	return db
}
