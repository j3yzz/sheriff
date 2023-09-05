package router

import (
	handler2 "github.com/j3yzz/sheriff/internal/delivery/http/handler"
	"github.com/j3yzz/sheriff/internal/infrastructure/config"
	"github.com/j3yzz/sheriff/internal/repository"
	"github.com/j3yzz/sheriff/internal/service/sms_service/kavenegarsvc"
	"github.com/labstack/echo/v4"
)

func Register(app *echo.Echo, repos *repository.Repositories, cfg config.Config) {
	kavenegarSvc := kavenegarsvc.New(cfg.SmsService)

	handler2.Health{}.Register(app.Group("/api"))

	handler2.Auth{
		Store:         repos.UserRepository,
		SmsService:    kavenegarSvc,
		OtpTokenStore: repos.OtpTokenRepository,
	}.Register(app.Group("/api"))
}
