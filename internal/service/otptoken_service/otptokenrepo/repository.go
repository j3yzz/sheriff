package otptokenrepo

import (
	"github.com/j3yzz/sheriff/internal/model"
	"github.com/j3yzz/sheriff/internal/service/otptoken_service/otptokenentity"
)

type Repository interface {
	CreateOtpToken(e otptokenentity.OtpTokenCreateEntity) (model.OtpToken, error)
	FindByToken(e otptokenentity.OtpTokenCreateEntity) (model.OtpToken, error)
	FindValidOTPToken(userID uint) (model.OtpToken, error)
	FindValidOTPTokenByUserAndToken(userID uint, token string) (model.OtpToken, error)
}
