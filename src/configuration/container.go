package configuration

import (
	"github.com/snkonoplev/file-manager/db"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()
	container.Provide(NewConfiguration)
	container.Provide(db.NewDbConnection)
	container.Provide(db.NewRepository)

	return container
}
