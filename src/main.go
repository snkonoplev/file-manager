package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/snkonoplev/file-manager/bootstrap"
	"github.com/snkonoplev/file-manager/configuration"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	logrus.Info("starting application")

	container := configuration.BuildContainer()
	err := container.Invoke(func(b *bootstrap.Bootstrap) error {
		err := b.Run()
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
}
