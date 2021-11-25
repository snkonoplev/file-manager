package configuration

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/snkonoplev/file-manager/auth"
	"github.com/snkonoplev/file-manager/bootstrap"
	"github.com/snkonoplev/file-manager/command"
	"github.com/snkonoplev/file-manager/commandhandler"
	"github.com/snkonoplev/file-manager/controller"
	"github.com/snkonoplev/file-manager/db"
	"github.com/snkonoplev/file-manager/mediator"
	"github.com/snkonoplev/file-manager/query"
	"github.com/snkonoplev/file-manager/queryhandler"
	"github.com/snkonoplev/file-manager/router"
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
	container.Provide(mediator.NewMediator)
	container.Provide(auth.NewAuth)
	container.Provide(bootstrap.NewSwagger)

	container.Provide(router.NewRouter)
	container.Provide(func() *gin.Engine {
		r := gin.New()
		r.Use(
			gin.Recovery(),
		)
		return r
	})

	registerHandlers(container)
	registerHttpHandlers(container)

	return container
}

func registerHttpHandlers(container *dig.Container) {
	container.Provide(controller.NewUsersController)
}

func registerHandlers(container *dig.Container) {

	container.Provide(queryhandler.NewUsersHandler)
	container.Provide(commandhandler.NewCreateUserHandler)
	container.Provide(queryhandler.NewAuthorizeHandler)
	container.Provide(commandhandler.NewUpdateUserHandler)
	container.Provide(commandhandler.NewDeleteUserHandler)

	container.Provide(func(usersHandler *queryhandler.UsersHandler) map[reflect.Type]mediator.Handler {
		return make(map[reflect.Type]mediator.Handler)
	})

	container.Invoke(func(handlers map[reflect.Type]mediator.Handler, handler *queryhandler.UsersHandler) {
		handlers[reflect.TypeOf(query.UsersQuery{})] = handler
	})
	container.Invoke(func(handlers map[reflect.Type]mediator.Handler, handler *commandhandler.CreateUserHandler) {
		handlers[reflect.TypeOf(command.CreateUserCommand{})] = handler
	})
	container.Invoke(func(handlers map[reflect.Type]mediator.Handler, handler *queryhandler.AuthorizeHandler) {
		handlers[reflect.TypeOf(query.UserAuthorizeQuery{})] = handler
	})
	container.Invoke(func(handlers map[reflect.Type]mediator.Handler, handler *commandhandler.UpdateUserHandler) {
		handlers[reflect.TypeOf(command.UpdateUserCommand{})] = handler
	})
	container.Invoke(func(handlers map[reflect.Type]mediator.Handler, handler *commandhandler.DeleteUserHandler) {
		handlers[reflect.TypeOf(command.DeleteUserCommand{})] = handler
	})
}
