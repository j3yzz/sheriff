package user_service

import (
	"github.com/j3yzz/sheriff/internal/model"
	"github.com/j3yzz/sheriff/internal/service/user_service/userentity"
	"github.com/j3yzz/sheriff/internal/service/user_service/userrepo"
)

type UserService struct {
	UserStore userrepo.Repository
}

func (u *UserService) Register(validatedEntity userentity.UserRegisterEntity) (model.User, error) {
	user, createUserErr := u.UserStore.CreateUser(validatedEntity)
	if createUserErr != nil {
		return model.User{}, createUserErr
	}

	return user, nil
}
