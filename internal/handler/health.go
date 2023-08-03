package handler

import (
	"github.com/j3yzz/sheriff/internal/pkg/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Health struct{}

func (h Health) Handle(c echo.Context) error {
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Data:    "everything is fine!",
	})
}

func (h Health) Register(g *echo.Group) {
	g.GET("/health", h.Handle)
}
