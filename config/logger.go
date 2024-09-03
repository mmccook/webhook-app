package config

import (
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"os"
)

func ConfigureLogger(viper *viper.Viper) zerolog.Logger {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	return logger
}
