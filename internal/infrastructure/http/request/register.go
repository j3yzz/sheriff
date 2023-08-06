package request

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type RegisterRequest struct {
	Phone string `json:"phone"`
}

func (r RegisterRequest) Validate() error {
	if err := validation.ValidateStruct(
		&r,
		validation.Field(&r.Phone, validation.Required),
	); err != nil {
		return fmt.Errorf("register request validation failed: %w", err)
	}

	return nil
}
