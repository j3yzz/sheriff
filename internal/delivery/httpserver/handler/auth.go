package handler

import (
	"github.com/j3yzz/sheriff/internal/adapter/sms_adapter/kavenegar_adapter"
	"github.com/j3yzz/sheriff/internal/delivery/httpserver/httprequest"
	"github.com/j3yzz/sheriff/internal/pkg/response"
	"github.com/j3yzz/sheriff/internal/service/otptoken_service/otptokenrepo"
	"github.com/j3yzz/sheriff/internal/service/user_service"
	"github.com/j3yzz/sheriff/internal/service/user_service/userrepo"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Auth struct {
	Store         userrepo.Repository
	SmsService    *kavenegar_adapter.KavenegarSvc
	OtpTokenStore otptokenrepo.Repository
}

func (a Auth) RegisterHandler(c echo.Context) error {
	var req httprequest.RegisterRequest
	validatedEntity, validatedErr := req.Validated(c)
	if validatedErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: validatedErr.Error(),
		})
	}

	svc := user_service.UserService{
		UserStore: a.Store,
	}

	user, err := svc.Register(validatedEntity)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Data: struct {
			ID    uint   `json:"id"`
			Phone string `json:"phone"`
		}{
			ID:    user.ID,
			Phone: user.Phone,
		},
	})
}

func (a Auth) Register(g *echo.Group) {
	g.POST("/auth/register", a.RegisterHandler)
}
