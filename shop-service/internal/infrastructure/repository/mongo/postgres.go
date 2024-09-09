package mongodb

import (
	"database/sql"

	"gitlab.com/acumen5524834/shop-service/internal/infrastructure/repository"
)

type Storage struct {
	OrderS repository.Orders
}

func NewOrderRepo(db *sql.DB) repository.Orders {
	return &OrderRepository{
		db: db,
	}
}

