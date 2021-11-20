package main

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	"github.com/golang-migrate/migrate/source/file"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "manager.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = RunMigrateScripts(db)
	if err != nil {
		panic(err)
	}
}

func RunMigrateScripts(db *sql.DB) error {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return fmt.Errorf("creating sqlite3 db driver failed %s", err)
	}

	fsrc, err := (&file.File{}).Open("file://migrations")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("file", fsrc, "manager.db", driver)
	if err != nil {
		return fmt.Errorf("initializing db migration failed %s", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migrating database failed %s", err)
	}

	return nil
}
