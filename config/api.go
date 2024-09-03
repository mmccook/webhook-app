package config

import (
	"DJMIL/api/template"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

type ApiConfig struct {
	DB       *redis.Client
	App      *echo.Echo
	Log      zerolog.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func ApiBootstrap(config *ApiConfig) {

}

func SetupApp(viper *viper.Viper) *echo.Echo {
	echoApp := echo.New()
	SetupEchoMiddleware(echoApp, viper)
	return echoApp
}

func SetupServer(config *ApiConfig) *http.Server {

	var server = http.Server{
		Addr:    fmt.Sprintf(":%d", config.Config.GetInt("API_PORT")),
		Handler: config.App,
	}
	return &server
}

func SetupEchoMiddleware(echoApp *echo.Echo, config *viper.Viper) {
	template.NewTemplateRenderer(echoApp)
	logger := zerolog.New(os.Stdout)

	echoApp.Pre(middleware.RemoveTrailingSlash())
	echoApp.Use(session.Middleware(sessions.NewCookieStore([]byte(config.GetString("SESSION_NAME")))))
	echoApp.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))
}
