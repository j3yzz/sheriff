package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Health struct {
}

type HandleResponse struct {
	Status  bool
	Message string
}

func (h Health) Handle(c echo.Context) error {
	return c.JSON(http.StatusOK, HandleResponse{
		Status:  true,
		Message: "everything is fine!",
	})
}

func (h Health) Register(g *echo.Group) {
	g.GET("/health", h.Handle)
}
