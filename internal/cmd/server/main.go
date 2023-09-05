package server

import (
	"fmt"
	"github.com/j3yzz/sheriff/internal/delivery/httpserver/server"
	"github.com/j3yzz/sheriff/internal/infrastructure/config"
	"github.com/j3yzz/sheriff/internal/infrastructure/db"
	"github.com/j3yzz/sheriff/internal/repository"
	"github.com/j3yzz/sheriff/internal/service/authservice"
	"github.com/j3yzz/sheriff/internal/service/user_service"
	"github.com/j3yzz/sheriff/internal/service/user_service/userentity"
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

				s := user_service.UserService{
					UserStore: repos.UserRepository,
					AuthSvc:   authservice.New(cfg.AuthService),
				}
				login, err := s.Login(userentity.LoginEntity{
					Phone:    "09123456789",
					Password: "123123123123",
				})
				fmt.Println("login", login)
				fmt.Println("err", err)

				server.Provide(repos, cfg)
			},
		},
	)
}
