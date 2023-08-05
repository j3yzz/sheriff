package server

import (
	"github.com/j3yzz/sheriff/internal/infrastructure/config"
	"github.com/j3yzz/sheriff/internal/infrastructure/db"
	"github.com/j3yzz/sheriff/internal/infrastructure/http/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func Register(root *cobra.Command) {
	root.AddCommand(
		&cobra.Command{
			Use:   "server",
			Short: "Run server",
			Run: func(_ *cobra.Command, _ []string) {
				fx.New(
					fx.Provide(config.Provide),
					db.Module,
					fx.Provide(server.Provide),
					fx.Invoke(func(e *echo.Echo) {
						log.Info("Web server module invoked")
					}, func(d *db.GormDatabase) {
						log.Info("Database module invoked")
					}),
				).Run()
			},
		},
	)
}
