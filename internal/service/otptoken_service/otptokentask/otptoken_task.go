package otptokentask

import (
	"github.com/j3yzz/sheriff/internal/model"
	"github.com/j3yzz/sheriff/internal/service/otptoken_service/otptokenentity"
	"github.com/j3yzz/sheriff/internal/service/otptoken_service/otptokenrepo"
	"math/rand"
	"strconv"
	"time"
)

type OtpTokenTask struct {
	Store otptokenrepo.Repository
}

func (t OtpTokenTask) CreateNewOtpToken(user model.User) (model.OtpToken, error) {
	var token int

	_, err := t.Store.FindValidOTPToken(user.ID)
	if err != nil {
		return model.OtpToken{}, err
	}

	for {
		token = 00000 + rand.Intn(99999)
		foundOtpToken, _ := t.Store.FindByToken(otptokenentity.OtpTokenCreateEntity{
			UserID: user.ID,
			Token:  strconv.Itoa(token),
		})

		if foundOtpToken.ID == 0 {
			break
		}
	}

	otpToken, err := t.Store.CreateOtpToken(otptokenentity.OtpTokenCreateEntity{
		UserID:   user.ID,
		Token:    strconv.Itoa(token),
		ExpireAt: time.Now().Add(time.Minute * 2),
	})

	if err != nil {
		return model.OtpToken{}, err
	}

	return otpToken, nil
}
