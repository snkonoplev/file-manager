package bootstrap

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/snkonoplev/file-manager/comand"
	"github.com/snkonoplev/file-manager/db"
	"github.com/spf13/viper"
)

type Bootstrap struct {
	config     *viper.Viper
	database   *sqlx.DB
	repository *db.Repository
}

func NewBootstrap(config *viper.Viper, database *sqlx.DB, repository *db.Repository) *Bootstrap {
	return &Bootstrap{
		config:     config,
		database:   database,
		repository: repository,
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
	exists, err := b.repository.CheckUserExists(context.Background(), "admin")
	if err != nil {
		return false, fmt.Errorf("can't check if user exists %s", err)
	}

	if exists {
		return false, nil
	}

	password := b.config.GetString("ADMIN_PASSWORD")
	_, err = b.repository.CreateUser(context.Background(), comand.CreateUserCommand{Name: "admin", Password: password, IsAdmin: true})
	if err != nil {
		return false, fmt.Errorf("can't create admin user %s", err)
	}
	return true, nil
}
