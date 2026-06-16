package suppliers

import (
	"database/sql"
	"errors"

	"github.com/brightDN/orderDesk/internal/database"
)

type SupplierService struct {
	queries *database.Queries
	db      *sql.DB
}

func NewSupplierService(queries *database.Queries, db *sql.DB) *SupplierService {
	return &SupplierService{
		queries: queries,
		db:      db,
	}
}

type Supplier struct {
	ID            int32
	Name          string
	Email         string
	ContactPerson string
	Count         int64
	Active        bool
	MailSubject   string
	MailContext   string
}

var ErrInternalError = errors.New("Something went wrong on our end, please try again later")
