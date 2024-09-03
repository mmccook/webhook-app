package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func ConfigureViper() *viper.Viper {
	config := viper.New()

	config.SetConfigName("app")
	config.SetConfigType("env")
	config.AddConfigPath("../../")
	config.AddConfigPath("./")
	err := config.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Failed setting up the config: %w \n", err))
	}

	return config
}
