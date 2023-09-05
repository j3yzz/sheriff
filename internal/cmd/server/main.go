package server

import (
	"github.com/j3yzz/sheriff/internal/infrastructure/config"
	"github.com/j3yzz/sheriff/internal/infrastructure/db"
	"github.com/j3yzz/sheriff/internal/infrastructure/http/server"
	"github.com/j3yzz/sheriff/internal/repository"
	"github.com/spf13/cobra"
)

func Register(root *cobra.Command) {
	root.AddCommand(
		&cobra.Command{
			Use:   "server",
			Short: "Run server",
			Run: func(_ *cobra.Command, _ []string) {
				cfg := config.Provide()
				database := db.New(cfg.Database)
				repos := repository.ProvideRepos(database)
				server.Provide(repos, cfg)
			},
		},
	)
}
