package httprequest

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/j3yzz/sheriff/internal/pkg/response"
	"github.com/j3yzz/sheriff/internal/service/user_service/userentity"
	"github.com/labstack/echo/v4"
	"net/http"
)

type LoginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (r LoginRequest) Validate() error {
	if err := validation.ValidateStruct(
		&r,
		validation.Field(&r.Phone, validation.Required),
		validation.Field(&r.Password, validation.Required),
	); err != nil {
		return fmt.Errorf("register request validation failed: %w", err)
	}

	return nil
}

func (r LoginRequest) Validated(c echo.Context) (userentity.LoginEntity, error) {

	if err := c.Bind(&r); err != nil {
		return userentity.LoginEntity{}, echo.NewHTTPError(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	if err := r.Validate(); err != nil {
		return userentity.LoginEntity{}, echo.NewHTTPError(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	u := userentity.LoginEntity{
		Phone:    r.Phone,
		Password: r.Password,
	}

	return u, nil
}
