package server

import (
	"errors"
	"github.com/j3yzz/sheriff/internal/delivery/httpserver/router"
	"github.com/j3yzz/sheriff/internal/infrastructure/config"
	"github.com/j3yzz/sheriff/internal/repository"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func Provide(repos *repository.Repositories, cfg config.Config) {
	app := echo.New()

	app.Use(echoprometheus.NewMiddleware("sheriff"))
	app.GET("/metrics", echoprometheus.NewHandler())

	router.Register(app, repos, cfg)

	if err := app.Start(":8080"); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("echo init failed", err)
	}
}
