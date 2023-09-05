package repository

import (
	"github.com/j3yzz/sheriff/internal/infrastructure/db"
	"github.com/j3yzz/sheriff/internal/service/user_service/userrepo"
)

type Repositories struct {
	UserRepository *userrepo.UserRepository
}

func ProvideRepos(db *db.GormDatabase) *Repositories {
	userRepository := userrepo.New(db)

	return &Repositories{
		UserRepository: userRepository,
	}
}
