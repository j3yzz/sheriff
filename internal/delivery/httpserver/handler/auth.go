package handler

import (
	"github.com/j3yzz/sheriff/internal/delivery/httpserver/httprequest"
	"github.com/j3yzz/sheriff/internal/pkg/response"
	"github.com/j3yzz/sheriff/internal/service/user_service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Auth struct {
	UserSvc user_service.UserService
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

	user, err := a.UserSvc.Register(validatedEntity)
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
	g.POST("/register", a.RegisterHandler)
}
