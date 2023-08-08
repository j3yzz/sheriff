package router

import (
	"github.com/j3yzz/sheriff/internal/infrastructure/http/handler"
	"github.com/j3yzz/sheriff/internal/repository"
	"github.com/j3yzz/sheriff/internal/service/sms_service/kavenegarsvc"
	"github.com/labstack/echo/v4"
)

func Register(app *echo.Echo, repos *repository.Repositories, smsSvcCfg kavenegarsvc.Config) {
	handler.Health{}.Register(app.Group("/api"))

	handler.Auth{
		Store: repos.UserRepository,
	}.Register(app.Group("/api"))
}
