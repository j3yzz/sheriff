package handler

import (
	"github.com/j3yzz/sheriff/internal/infrastructure/http/request"
	"github.com/j3yzz/sheriff/internal/pkg/response"
	"github.com/j3yzz/sheriff/internal/service/sms_service/kavenegarsvc"
	"github.com/j3yzz/sheriff/internal/service/user_service/userentity"
	"github.com/j3yzz/sheriff/internal/service/user_service/userrepo"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Auth struct {
	Store      userrepo.Repository
	SmsService *kavenegarsvc.KavenegarSvc
}

func (a Auth) RegisterHandler(c echo.Context) error {
	var req request.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	if err := req.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	u := userentity.UserRegisterEntity{
		Phone: req.Phone,
	}

	user, err := a.Store.CreateUser(u)
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

func (a Auth) RequestTokenHandler(c echo.Context) error {
	var req request.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	if err := req.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	// find phone number
	_, err := a.Store.FindByPhone(req.Phone)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, response.ErrorResponse{
			Success: false,
			Message: "user.not_found",
		})
	}

	// create otp token
	// send otp token to user

	return c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Data:    "otp.sent",
	})
}

func (a Auth) Register(g *echo.Group) {
	g.POST("/auth/register", a.RegisterHandler)
	g.POST("/auth/token/request", a.RequestTokenHandler)
}
