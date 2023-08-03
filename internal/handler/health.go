package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Health struct {
}

type HealthResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func (h Health) Handle(c echo.Context) error {
	return c.JSON(http.StatusOK, HealthResponse{
		Status:  true,
		Message: "everything is fine!",
	})
}

func (h Health) Register(g *echo.Group) {
	g.GET("/health", h.Handle)
}
