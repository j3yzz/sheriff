package handler

import (
	"github.com/j3yzz/sheriff/internal/infrastructure/http/request"
	"github.com/j3yzz/sheriff/internal/pkg/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Auth struct{}

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

	return c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Data:    req.Phone,
	})
}

func (a Auth) Register(g *echo.Group) {
	g.POST("/auth/register", a.RegisterHandler)
}
