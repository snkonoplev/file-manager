package configuration

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func NewConfiguration() (*viper.Viper, error) {
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
		return nil, fmt.Errorf("can't read config file %w", err)
	}

	v.AutomaticEnv()

	return v, nil
}
