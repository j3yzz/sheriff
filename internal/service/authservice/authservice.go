package authservice

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/j3yzz/sheriff/internal/service/authservice/accesstokenrepo"
	"github.com/j3yzz/sheriff/internal/service/user_service/model"
	"strings"
	"time"
)

type AuthService struct {
	config          Config
	accessTokenRepo accesstokenrepo.Repository
}

func New(config Config, accessTokenRepo accesstokenrepo.Repository) AuthService {
	return AuthService{
		config:          config,
		accessTokenRepo: accessTokenRepo,
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

func (s AuthService) ParseToken(tokenString string) (*Claims, error) {
	tokenString = strings.Replace(tokenString, "Bearer ", "", -1)

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.SignKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
