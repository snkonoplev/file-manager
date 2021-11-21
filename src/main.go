package main

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/snkonoplev/file-manager/bootstrap"
	"github.com/snkonoplev/file-manager/configuration"
	"github.com/snkonoplev/file-manager/mediator"
	"github.com/snkonoplev/file-manager/query"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	logrus.Info("starting application")

	container := configuration.BuildContainer()
	err := container.Invoke(func(b *bootstrap.Bootstrap, mediator *mediator.Mediator) error {
		err := b.Run()
		if err != nil {
			return err
		}

		users, err := mediator.Handle(context.Background(), query.UsersQuery{})
		if err != nil {
			return err
		}

		logrus.WithField("users", users).Info("users list")

		return nil
	})
	if err != nil {
		panic(err)
	}
}
