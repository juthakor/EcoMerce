package domain

import "time"

type Order struct {
	ID         string
	UserID     string
	ProductIDs []string
	Total      float64
	Status     string
	CreatedAt  time.Time
}
