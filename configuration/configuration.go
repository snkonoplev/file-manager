package configuration

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewConfiguration(logger *logrus.Logger) *viper.Viper {
	v := viper.NewWithOptions(viper.KeyDelimiter("_"))
	v.SetConfigName("appsettings")
	v.SetConfigType("yml")

	if path, ok := os.LookupEnv("CONFIG_FILE_PATH"); ok {
		v.AddConfigPath(path)
	} else {
		v.AddConfigPath(".")
	}
	err := v.ReadInConfig()

	if err != nil {
		logger.Fatal(err)
	}

	v.AutomaticEnv()

	return v
}
