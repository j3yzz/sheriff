package server

import (
	"context"
	"errors"
	"github.com/j3yzz/sheriff/internal/infrastructure/router"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"log"
	"net/http"
)

func Provide(lc fx.Lifecycle) *echo.Echo {
	app := echo.New()
	router.Register(app)

	lc.Append(
		fx.Hook{
			OnStart: func(_ context.Context) error {
				go func() {
					if err := app.Start(":8080"); !errors.Is(err, http.ErrServerClosed) {
						log.Fatal("echo init failed", err)
					}
				}()
				return nil
			},
			OnStop: app.Shutdown,
		},
	)

	return app
}
