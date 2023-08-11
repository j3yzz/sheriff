package userrepo

import (
	"github.com/j3yzz/sheriff/internal/model"
	"github.com/j3yzz/sheriff/internal/service/user_service/userentity"
)

type Repository interface {
	CreateUser(user userentity.UserRegisterEntity) (model.User, error)
	FindByPhone(phone string) (model.User, error)
}
