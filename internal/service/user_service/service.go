package user_service

import (
	"github.com/j3yzz/sheriff/internal/service/user_service/model"
	"github.com/j3yzz/sheriff/internal/service/user_service/userentity"
	"github.com/j3yzz/sheriff/internal/service/user_service/userrepo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserStore userrepo.Repository
}

func (u *UserService) Register(validatedEntity userentity.UserRegisterEntity) (model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(validatedEntity.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return model.User{}, err
	}

	validatedEntity.Password = string(hashedPassword)

	user, createUserErr := u.UserStore.CreateUser(validatedEntity)
	if createUserErr != nil {
		return model.User{}, createUserErr
	}

	return user, nil
}
