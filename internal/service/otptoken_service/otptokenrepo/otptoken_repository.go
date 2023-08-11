package otptokenrepo

import (
	"errors"
	"github.com/j3yzz/sheriff/internal/infrastructure/db"
	"github.com/j3yzz/sheriff/internal/model"
	"github.com/j3yzz/sheriff/internal/service/otptoken_service/otptokenentity"
	"time"
)

type OtpTokenRepository struct {
	db *db.GormDatabase
}

func New(db *db.GormDatabase) *OtpTokenRepository {
	return &OtpTokenRepository{
		db: db,
	}
}

const tableName = "otp_tokens"
const FoundValidOTPToken = "find.valid.otp_token"

func (r *OtpTokenRepository) CreateOtpToken(e otptokenentity.OtpTokenCreateEntity) (model.OtpToken, error) {
	result := r.db.DB.Table(tableName).Create(&e)

	if result.Error != nil {
		return model.OtpToken{}, result.Error
	}

	return model.OtpToken{
		ID:       e.ID,
		Token:    e.Token,
		UserID:   e.UserID,
		ExpireAt: e.ExpireAt,
	}, nil
}

func (r *OtpTokenRepository) FindByToken(e otptokenentity.OtpTokenCreateEntity) (model.OtpToken, error) {
	var otpToken model.OtpToken
	result := r.db.DB.Table(tableName).Where("token = ? and user_id = ?", e.Token, e.UserID).First(&otpToken)

	if result.Error != nil {
		return model.OtpToken{}, result.Error
	}

	return otpToken, nil
}

func (r *OtpTokenRepository) FindValidOTPToken(userID uint) (model.OtpToken, error) {
	var otpToken model.OtpToken
	result := r.db.DB.Table(tableName).Where("user_id = ? and expire_at > ?", userID, time.Now()).Last(&otpToken)

	if result.Error != nil {
		return model.OtpToken{}, result.Error
	}

	if result.RowsAffected != 0 {
		return model.OtpToken{}, errors.New(FoundValidOTPToken)
	}

	return otpToken, nil
}
