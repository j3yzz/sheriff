package repository

import (
	"github.com/j3yzz/sheriff/internal/infrastructure/db"
	"github.com/j3yzz/sheriff/internal/service/otptoken_service/otptokenrepo"
	"github.com/j3yzz/sheriff/internal/service/user_service/userrepo"
)

type Repositories struct {
	UserRepository     *userrepo.UserRepository
	OtpTokenRepository *otptokenrepo.OtpTokenRepository
}

func ProvideRepos(db *db.GormDatabase) *Repositories {
	userRepository := userrepo.New(db)
	otpTokenRepository := otptokenrepo.New(db)

	return &Repositories{
		UserRepository:     userRepository,
		OtpTokenRepository: otpTokenRepository,
	}
}
