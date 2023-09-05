package userrepo

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/j3yzz/sheriff/internal/infrastructure/db"
	"github.com/j3yzz/sheriff/internal/service/user_service/model"
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

const tableName = "users"

func (r *UserRepository) CreateUser(user userentity.UserRegisterEntity) (model.User, error) {
	var mysqlErr *mysql.MySQLError
	result := r.db.DB.Table(tableName).Omit("id").Create(&user)

	if errors.As(result.Error, &mysqlErr) && mysqlErr.Number == 1062 {
		return model.User{}, errors.New("user.already.exists")
	}

	if result.Error != nil {
		return model.User{}, result.Error
	}

	return model.User{
		ID:    user.ID,
		Phone: user.Phone,
	}, nil
}

func (r *UserRepository) FindByPhone(phone string) (model.User, error) {
	var user model.User
	result := r.db.DB.Table(tableName).Where("phone = ?", phone).First(&user)

	if result.Error != nil {
		return model.User{}, result.Error
	}

	if result.RowsAffected == 0 {
		return model.User{}, errors.New("user not found")
	}

	return user, nil
}
