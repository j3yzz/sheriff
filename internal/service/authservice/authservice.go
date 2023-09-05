package authservice

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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
	tokenPayload := AccessTokenSubject + "/" + uuid.New().String()
	expirationTime := time.Now().Add(s.config.AccessExpirationTime)

	token, err := createToken(user, expirationTime, s.config.SignKey, AccessTokenSubject, tokenPayload)
	if err != nil {
		return "", err
	}

	_, err = s.accessTokenRepo.CreateAccessToken(tokenPayload, user.ID, expirationTime, "127.0.05.2", "firefox")
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s AuthService) CreateRefreshToken(user model.User) (string, error) {
	tokenPayload := RefreshTokenSubject + "/" + uuid.New().String()
	expirationTime := time.Now().Add(s.config.RefreshExpirationTime)

	token, err := createToken(user, expirationTime, s.config.SignKey, RefreshTokenSubject, tokenPayload)
	if err != nil {
		return "", err
	}

	_, err = s.accessTokenRepo.CreateAccessToken(tokenPayload, user.ID, expirationTime, "127.0.05.2", "firefox")
	if err != nil {
		return "", err
	}

	return token, nil
}

func createToken(user model.User, expirationTime time.Time, signKey string, tokenSubject string, tokenPayload string) (string, error) {
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   tokenSubject,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
		UserID: user.ID,
		Token:  tokenPayload,
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
