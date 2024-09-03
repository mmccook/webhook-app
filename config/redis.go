package config

import (
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

func ConfigureRedisDB(viper *viper.Viper, log zerolog.Logger) (*redis.Client, error) {
	var host string = viper.GetString("REDIS_HOST")
	if host == "" {
		return nil, errors.New("REDIS_HOST must be set")
	}

	var port string = viper.GetString("REDIS_PORT")
	if port == "" {
		return nil, errors.New("REDIS_PORT must be set")
	}

	var db int = viper.GetInt("REDIS_DB")

	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: "",
		DB:       db,
	}), nil
}
