package server

import (
	"github.com/j3yzz/sheriff/internal/infrastructure/config"
	"github.com/j3yzz/sheriff/internal/infrastructure/db"
	"github.com/j3yzz/sheriff/internal/infrastructure/http/server"
	"github.com/j3yzz/sheriff/internal/repository"
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
					fx.Provide(repository.ProvideRepos),
					fx.Invoke(func(e *echo.Echo) {
						log.Info("Web server module invoked")
					}, func(d *db.GormDatabase) {
						log.Info("Database module invoked")
					}, func(migrator db.MigrateError) {
						if err := migrator(); err != nil {
							log.Fatalf("error in migrating database: %v", err)
						} else {
							log.Info("Migrating database invoked")
						}
					}, func(repo *repository.Repositories) {
						log.Info("Repositories invoked")
					}),
				).Run()
			},
		},
	)
}
