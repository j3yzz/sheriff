package authservice

import "time"

const (
	AccessTokenSubject  = "at"
	RefreshTokenSubject = "rt"
)

type Config struct {
	SignKey               string        `json:"sign_key" koanf:"sign_key"`
	AccessExpirationTime  time.Duration `json:"access_expiration_time" koanf:"access_expiration_time"`
	RefreshExpirationTime time.Duration `json:"refresh_expiration_time" koanf:"refresh_expiration_time"`
}
