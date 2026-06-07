package suppliers

import (
	"errors"

	"github.com/brightDN/orderDesk/internal/database"
)

type SupplierService struct {
	db *database.Queries
}

func NewSupplierService(db *database.Queries) *SupplierService {
	return &SupplierService{db: db}
}

type Supplier struct {
	ID            int32
	Name          string
	Email         string
	ContactPerson string
	Count         int64
	Active        bool
}

var ErrInternalError = errors.New("Something went wrong on our end, please try again later")
