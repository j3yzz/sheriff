package accesstokenrepo

import (
	"github.com/j3yzz/sheriff/internal/infrastructure/db"
	"github.com/j3yzz/sheriff/internal/service/authservice/model"
	"time"
)

type AccessTokenRepository struct {
	db *db.GormDatabase
}

func New(db *db.GormDatabase) *AccessTokenRepository {
	return &AccessTokenRepository{
		db: db,
	}
}

const tableName = "access_tokens"

func (a *AccessTokenRepository) CreateAccessToken(token string, userID uint, expireTime time.Time, ipAddress string, userAgent string) (model.AccessToken, error) {
	accessToken := model.AccessToken{
		UserID:     userID,
		IpAddress:  ipAddress,
		UserAgent:  userAgent,
		Token:      token,
		ExpireTime: expireTime,
	}

	result := a.db.DB.Table(tableName).Omit("id").Create(&accessToken)

	if result.Error != nil {
		return model.AccessToken{}, result.Error
	}

	return accessToken, nil
}

func (a *AccessTokenRepository) FindByToken(token string) (model.AccessToken, error) {
	panic("implement me")
}
