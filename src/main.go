package main

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/snkonoplev/file-manager/comand"
	"github.com/snkonoplev/file-manager/configuration"
	"github.com/snkonoplev/file-manager/db"
)

func main() {
	container := configuration.BuildContainer()
	err := container.Invoke(func(database *sqlx.DB, repository *db.Repository) {
		db.RunMigrateScripts(database.DB)
		repository.CreateUser(context.Background(), comand.CreateUserCommand{Name: "admin", Password: "123"})

	})
	if err != nil {
		panic(err)
	}
}
