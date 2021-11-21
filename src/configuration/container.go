package configuration

import (
	"reflect"

	"github.com/jmoiron/sqlx"
	"github.com/snkonoplev/file-manager/bootstrap"
	"github.com/snkonoplev/file-manager/command"
	"github.com/snkonoplev/file-manager/commandhandler"
	"github.com/snkonoplev/file-manager/db"
	"github.com/snkonoplev/file-manager/mediator"
	"github.com/snkonoplev/file-manager/query"
	"github.com/snkonoplev/file-manager/queryhandler"
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
	container.Provide(db.NewUsersRepository)
	container.Provide(bootstrap.NewBootstrap)
	registerHandlers(container)
	container.Provide(mediator.NewMediator)

	return container
}

func registerHandlers(container *dig.Container) {

	container.Provide(queryhandler.NewUsersHandler)
	container.Provide(commandhandler.NewCreateUserHandler)

	container.Provide(func(usersHandler *queryhandler.UsersHandler) map[reflect.Type]mediator.Handler {
		return make(map[reflect.Type]mediator.Handler)
	})

	container.Invoke(func(handlers map[reflect.Type]mediator.Handler, handler *queryhandler.UsersHandler) {
		handlers[reflect.TypeOf(query.UsersQuery{})] = handler
	})
	container.Invoke(func(handlers map[reflect.Type]mediator.Handler, handler *commandhandler.CreateUserHandler) {
		handlers[reflect.TypeOf(command.CreateUserCommand{})] = handler
	})
}
