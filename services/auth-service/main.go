package main

import (
	"github.com/juthakor/EcoMerce/services/auth-service/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Global middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route group setup
	router.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":8081"))
}
