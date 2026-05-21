package companies

import (
	"errors"

	"github.com/brightDN/orderDesk/internal/database"
)

type CompanyService struct {
	db *database.Queries
}

func NewCompanyService(db *database.Queries) *CompanyService {
	return &CompanyService{
		db: db,
	}
}

// ERROR DECLARATIONS

var ErrInternalError = errors.New("something went wrong on our end, try again later")
