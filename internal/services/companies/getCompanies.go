package companies

import (
	"github.com/brightDN/orderDesk/internal/database"
	"github.com/labstack/echo/v4"
)

func GetCompanies(db *database.Queries, c echo.Context) ([]Company, error) {
	companiesData, err := db.GetCompanies(c.Request().Context())
	if err != nil {
		return nil, ErrInternalError
	}

	companies := make([]Company, 0, len(companiesData))

	for i, company := range companiesData {
		companies[i] = Company{
			ID:        company.ID,
			Name:      company.Name,
			Email:     company.Email,
			CreatedAt: company.CreatedAt.Format("02-01-2006"),
			IsDeleted: company.DeletedAt.Valid,
		}
	}

	return companies, nil
}
