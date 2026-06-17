package companies

import (
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
