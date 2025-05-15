package router

import (
	"github.com/juthakor/EcoMerce/services/auth-service/handler"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "auth-service OK")
	})

	// Auth routes
	e.POST("/register", handler.Register)
	e.POST("/login", handler.Login)
}
