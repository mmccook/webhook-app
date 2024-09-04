package main

import (
	"DJMIL/api/controller"
	"DJMIL/config"
	"fmt"
)

func main() {
	viperConfig := config.ConfigureViper()
	logger := config.ConfigureLogger(viperConfig)
	validator := config.ConfigureValidator(viperConfig)
	db, _ := config.ConfigureRedisDB(viperConfig, logger)
	app := config.SetupApp(viperConfig)

	appConfig := config.ApiConfig{
		DB:       db,
		Log:      logger,
		Validate: validator,
		App:      app,
		Config:   viperConfig,
	}

	config.ApiBootstrap(&appConfig)

	controller.InitSectionRouter(&appConfig)
	controller.InitWebhooksRouter(&appConfig)

	var server = config.SetupServer(&appConfig)

	err := server.ListenAndServe()
	if err != nil {
		logger.Fatal().Msg(fmt.Sprintf("Failed to start server: %v", err))
	}
}
