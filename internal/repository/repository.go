package repository

import (
	"github.com/j3yzz/sheriff/internal/infrastructure/db"
	"github.com/j3yzz/sheriff/internal/service/authservice/accesstokenrepo"
	"github.com/j3yzz/sheriff/internal/service/user_service/userrepo"
)

type Repositories struct {
	UserRepository        *userrepo.UserRepository
	AccessTokenRepository *accesstokenrepo.AccessTokenRepository
}

func ProvideRepos(db *db.GormDatabase) *Repositories {
	userRepository := userrepo.New(db)
	accessTokenRepository := accesstokenrepo.New(db)

	return &Repositories{
		UserRepository:        userRepository,
		AccessTokenRepository: accessTokenRepository,
	}
}
