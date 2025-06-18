package postgres

import (
	"database/sql"

	"github.com/juthakor/EcoMerce/services/order-service/internal/domain"
	"github.com/juthakor/EcoMerce/services/order-service/internal/repository"
)

type orderRepo struct {
	db *sql.DB
}

func NewPostgresOrderRepo(db *sql.DB) repository.OrderRepository {
	return &orderRepo{db: db}
}

func (r *orderRepo) FetchAll() ([]domain.Order, error) {
	rows, err := r.db.Query("SELECT id, user_id, total, status, created_at FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		var o domain.Order
		if err := rows.Scan(&o.ID, &o.UserID, &o.Total, &o.Status, &o.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}
