package server

import (
	"github.com/j3yzz/sheriff/internal/infrastructure/http/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func main(_ *echo.Echo) {
	log.Info("welcome to our server")
}

func Register(root *cobra.Command) {
	root.AddCommand(
		&cobra.Command{
			Use:   "server",
			Short: "Run server",
			Run: func(cmd *cobra.Command, _ []string) {
				fx.New(
					fx.Provide(server.Provide),
					fx.Invoke(main),
				).Run()
			},
		},
	)
}
