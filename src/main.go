package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/snkonoplev/file-manager/bootstrap"
	"github.com/snkonoplev/file-manager/configuration"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	logrus.Info("starting application")

	container := configuration.BuildContainer()
	err := container.Invoke(func(b *bootstrap.Bootstrap, engine *gin.Engine) error {
		err := b.Run()
		if err != nil {
			return err
		}

		err = engine.Run()
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		logrus.StandardLogger().Fatal(err)
	}

}
