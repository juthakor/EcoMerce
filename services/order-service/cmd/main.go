package main

import (
	"github.com/juthakor/EcoMerce/services/order-service/internal/delivery/http"
	"github.com/juthakor/EcoMerce/services/order-service/internal/infra"
	"github.com/juthakor/EcoMerce/services/order-service/internal/repository/postgres"
	"github.com/juthakor/EcoMerce/services/order-service/internal/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := infra.InitPostgres()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := postgres.NewPostgresOrderRepo(db)
	uc := usecase.NewOrderUsecase(repo)
	handler := http.NewHandler(uc)

	e := echo.New()
	handler.RegisterRoutes(e)

	e.Start(":8081")
}
