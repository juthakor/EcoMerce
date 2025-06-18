package http

import (
	"net/http"

	"github.com/juthakor/EcoMerce/services/order-service/internal/usecase"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Uc usecase.OrderUsecase
}

func NewHandler(uc usecase.OrderUsecase) *Handler {
	return &Handler{Uc: uc}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	e.GET("/orders", h.GetAll)
}

func (h *Handler) GetAll(c echo.Context) error {
	orders, err := h.Uc.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, orders)
}
