package handler

import (
	"github.com/j3yzz/sheriff/internal/infrastructure/http/request"
	"github.com/j3yzz/sheriff/internal/pkg/response"
	"github.com/j3yzz/sheriff/internal/service/user_service/userentity"
	"github.com/j3yzz/sheriff/internal/service/user_service/userrepo"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Auth struct {
	Store userrepo.Repository
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

func (a Auth) Register(g *echo.Group) {
	g.POST("/auth/register", a.RegisterHandler)
}
