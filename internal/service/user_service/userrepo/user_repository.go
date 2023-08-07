package userrepo

import (
	"github.com/j3yzz/sheriff/internal/infrastructure/db"
	"github.com/j3yzz/sheriff/internal/model"
	"github.com/j3yzz/sheriff/internal/service/user_service/userentity"
)

type UserRepository struct {
	db *db.GormDatabase
}

func New(db *db.GormDatabase) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user userentity.UserRegisterEntity) (model.User, error) {
	result := r.db.DB.Table("users").Omit("id").Create(&user)

	if result.Error != nil {
		return model.User{}, result.Error
	}

	return model.User{
		ID:    user.ID,
		Phone: user.Phone,
	}, nil
}
