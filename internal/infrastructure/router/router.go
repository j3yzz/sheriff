package router

import (
	"github.com/j3yzz/sheriff/internal/infrastructure/http/handler"
	"github.com/labstack/echo/v4"
)

func Register(app *echo.Echo) {
	handler.Health{}.Register(app.Group("/api"))
}
