package request

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/j3yzz/sheriff/internal/pkg/response"
	"github.com/j3yzz/sheriff/internal/service/user_service/userentity"
	"github.com/labstack/echo/v4"
	"net/http"
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

func (r RegisterRequest) Validated(c echo.Context) (userentity.UserRegisterEntity, error) {
	if err := c.Bind(&r); err != nil {
		return userentity.UserRegisterEntity{}, echo.NewHTTPError(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	if err := r.Validate(); err != nil {
		return userentity.UserRegisterEntity{}, echo.NewHTTPError(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	u := userentity.UserRegisterEntity{
		Phone: r.Phone,
	}

	return u, nil
}
