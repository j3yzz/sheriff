package user_service

import (
	"errors"
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

func (u *UserService) Login(e userentity.LoginEntity) (model.User, error) {
	// check user is exists or not and get user
	user, err := u.UserStore.FindByPhone(e.Phone)
	if err != nil {
		return model.User{}, err
	}

	// check password is compare with hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(e.Password))
	if err != nil {
		return model.User{}, errors.New("username or password isn't correct")
	}

	return user, err

	// create access token and refresh token

	// return user and tokens
}
