package usecase

import (
	"github.com/juthakor/EcoMerce/services/order-service/internal/domain"
	"github.com/juthakor/EcoMerce/services/order-service/internal/repository"
)

type OrderUsecase interface {
	GetAll() ([]domain.Order, error)
}

type orderUsecase struct {
	repo repository.OrderRepository
}

func NewOrderUsecase(r repository.OrderRepository) OrderUsecase {
	return &orderUsecase{repo: r}
}

func (uc *orderUsecase) GetAll() ([]domain.Order, error) {
	return uc.repo.FetchAll()
}
