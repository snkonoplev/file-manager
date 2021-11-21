package configuration

import (
	"github.com/jmoiron/sqlx"
	"github.com/snkonoplev/file-manager/bootstrap"
	"github.com/snkonoplev/file-manager/db"
	"github.com/spf13/viper"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()
	container.Provide(NewConfiguration)
	container.Provide(func(config *viper.Viper) (*sqlx.DB, error) {
		filepath := config.GetString("DB_FILE_PATH")
		return db.NewDbConnection(filepath)
	})
	container.Provide(db.NewRepository)
	container.Provide(bootstrap.NewBootstrap)

	return container
}
