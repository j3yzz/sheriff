package user_service

import (
	"errors"
	"github.com/j3yzz/sheriff/internal/service/authservice"
	"github.com/j3yzz/sheriff/internal/service/user_service/model"
	"github.com/j3yzz/sheriff/internal/service/user_service/userentity"
	"github.com/j3yzz/sheriff/internal/service/user_service/userrepo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserStore userrepo.Repository
	AuthSvc   authservice.AuthService
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

type UserInfo struct {
	ID    uint   `json:"ID"`
	Phone string `json:"phone"`
	Name  string `json:"name"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginResponse struct {
	User   UserInfo `json:"user"`
	Tokens Tokens   `json:"tokens"`
}

func (u *UserService) Login(e userentity.LoginEntity) (LoginResponse, error) {
	user, err := u.UserStore.FindByPhone(e.Phone)
	if err != nil {
		return LoginResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(e.Password))
	if err != nil {
		return LoginResponse{}, errors.New("username or password isn't correct")
	}

	accessToken, err := u.AuthSvc.CreateAccessToken(user)
	if err != nil {
		return LoginResponse{}, err
	}

	refreshToken, err := u.AuthSvc.CreateRefreshToken(user)
	if err != nil {
		return LoginResponse{}, err
	}

	return LoginResponse{
		User: UserInfo{
			ID:    user.ID,
			Phone: user.Phone,
			Name:  user.Name,
		},
		Tokens: Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}
