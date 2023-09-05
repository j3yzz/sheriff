package httprequest

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type VerifyTokenRequest struct {
	Phone string `json:"phone"`
	Token string `json:"token"`
}

func (r VerifyTokenRequest) Validate() error {
	if err := validation.ValidateStruct(
		&r,
		validation.Field(&r.Phone, validation.Required),
		validation.Field(&r.Token, validation.Required),
	); err != nil {
		return fmt.Errorf("verify token request validation failed: %w", err)
	}

	return nil
}
