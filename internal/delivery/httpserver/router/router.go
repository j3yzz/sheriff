package router

import (
	"github.com/j3yzz/sheriff/internal/delivery/httpserver/handler"
	"github.com/j3yzz/sheriff/internal/infrastructure/config"
	"github.com/j3yzz/sheriff/internal/repository"
	"github.com/labstack/echo/v4"
)

func Register(app *echo.Echo, repos *repository.Repositories, cfg config.Config) {
	apiApp := app.Group("/api")

	handler.Health{}.Register(apiApp)

	handler.Auth{
		Store:         repos.UserRepository,
		OtpTokenStore: repos.OtpTokenRepository,
	}.Register(apiApp.Group("/auth"))
}
