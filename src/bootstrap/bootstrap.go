package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/snkonoplev/file-manager/command"
	"github.com/snkonoplev/file-manager/db"
	"github.com/snkonoplev/file-manager/mediator"
	"github.com/spf13/viper"
)

type Bootstrap struct {
	config   *viper.Viper
	database *sqlx.DB
	mediator *mediator.Mediator
}

func NewBootstrap(config *viper.Viper, database *sqlx.DB, mediator *mediator.Mediator) *Bootstrap {
	return &Bootstrap{
		config:   config,
		database: database,
		mediator: mediator,
	}
}

func (b *Bootstrap) Run() error {

	err := b.runMigrations()
	if err != nil {
		return fmt.Errorf("can't migrate db %s", err)
	}
	logrus.Info("migrations compleated")

	created, err := b.createAdminUser()
	if err != nil {
		return err
	}

	if created {
		logrus.Info("admin user created")
	} else {
		logrus.Info("admin user already exists")
	}

	return nil
}

func (b *Bootstrap) runMigrations() error {
	filepath := b.config.GetString("DB_FILE_PATH")
	err := db.RunMigrateScripts(b.database.DB, filepath)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bootstrap) createAdminUser() (bool, error) {
	_, err := b.mediator.Handle(context.Background(), command.CreateUserCommand{

		Name:     "admin",
		Password: b.config.GetString("ADMIN_PASSWORD"),
		IsAdmin:  true,
	})
	if err != nil {

		target := &mediator.HandlerError{}
		if errors.As(err, &target) && target.StatusCode == http.StatusBadRequest {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
