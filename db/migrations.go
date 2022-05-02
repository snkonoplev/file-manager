package db

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	"github.com/golang-migrate/migrate/source/file"
)

func RunMigrateScripts(db *sql.DB, filepath string) error {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return fmt.Errorf("creating sqlite3 db driver failed %w", err)
	}

	f, err := (&file.File{}).Open("file://../migrations")
	if err != nil {
		return fmt.Errorf("can't get migration files %w", err)
	}

	m, err := migrate.NewWithInstance("file", f, filepath, driver)
	if err != nil {
		return fmt.Errorf("initializing db migration failed %w", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migrating database failed %w", err)
	}

	return nil
}
