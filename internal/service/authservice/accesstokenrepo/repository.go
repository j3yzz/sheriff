package accesstokenrepo

import (
	"github.com/j3yzz/sheriff/internal/service/authservice/model"
	"time"
)

type Repository interface {
	CreateAccessToken(token string, userID uint, expireTime time.Time, ipAddress string, userAgent string) (model.AccessToken, error)
	FindByToken(token string) (model.AccessToken, error)
}
