package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c echo.Context) error {
	var req UserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	// TODO: hash password, store to DB
	return c.JSON(http.StatusCreated, echo.Map{"message": "User registered (stub)"})
}

func Login(c echo.Context) error {
	var req UserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	// TODO: validate user, issue JWT
	return c.JSON(http.StatusOK, echo.Map{"token": "dummy.jwt.token"})
}
