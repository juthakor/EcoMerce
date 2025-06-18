package repository

import "github.com/juthakor/EcoMerce/services/order-service/internal/domain"

type OrderRepository interface {
	FetchAll() ([]domain.Order, error)
}
