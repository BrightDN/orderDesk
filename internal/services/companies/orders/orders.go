package orders

import (
	"database/sql"

	"github.com/brightDN/orderDesk/internal/database"
)

type OrderService struct {
	queries *database.Queries
	db      *sql.DB
}

func NewOrderService(queries *database.Queries, db *sql.DB) *OrderService {
	return &OrderService{
		queries: queries,
		db:      db,
	}
}
