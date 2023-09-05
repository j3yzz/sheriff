package router

import (
	"github.com/j3yzz/sheriff/internal/delivery/httpserver/handler"
	"github.com/j3yzz/sheriff/internal/infrastructure/config"
	"github.com/j3yzz/sheriff/internal/repository"
	"github.com/j3yzz/sheriff/internal/service/authservice"
	"github.com/j3yzz/sheriff/internal/service/user_service"
	"github.com/labstack/echo/v4"
)

func Register(app *echo.Echo, repos *repository.Repositories, cfg config.Config) {
	apiApp := app.Group("/api")

	handler.Health{}.Register(apiApp)

	authSvc := authservice.New(cfg.AuthService)

	userSvc := user_service.UserService{
		UserStore: repos.UserRepository,
		AuthSvc:   authSvc,
	}

	handler.Auth{
		UserSvc: userSvc,
	}.Register(apiApp.Group("/auth"))
}
