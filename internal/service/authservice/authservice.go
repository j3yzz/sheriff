package authservice

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/j3yzz/sheriff/internal/service/user_service/model"
	"time"
)

type AuthService struct {
	config Config
}

func New(config Config) AuthService {
	return AuthService{
		config: config,
	}
}

func (s AuthService) CreateAccessToken(user model.User) (string, error) {
	return createToken(user, s.config.AccessExpirationTime, s.config.SignKey, AccessTokenSubject)
}

func (s AuthService) CreateRefreshToken(user model.User) (string, error) {
	return createToken(user, s.config.RefreshExpirationTime, s.config.SignKey, RefreshTokenSubject)
}

func createToken(user model.User, expirationTime time.Duration, signKey string, tokenSubject string) (string, error) {
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   tokenSubject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
		},
		UserID: user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(signKey))
}
