package migration

import (
	"github.com/j3yzz/sheriff/internal/infrastructure/config"
	"github.com/j3yzz/sheriff/internal/infrastructure/db"
	"github.com/spf13/cobra"
	"log"
)

func Register(root *cobra.Command) {
	root.AddCommand(
		&cobra.Command{
			Use:   "migrate",
			Short: "run migration files",
			Run: func(_ *cobra.Command, _ []string) {
				cfg := config.Provide()
				database := db.New(cfg.Database)

				err := db.RunMigrations(database)
				if err != nil {
					log.Fatal(err)
				}
			},
		},
	)
}
