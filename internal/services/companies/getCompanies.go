package companies

import (
	"github.com/labstack/echo/v4"
)

func (cs *CompanyService) GetCompanies(c echo.Context) ([]Company, error) {
	companiesData, err := cs.db.GetCompanies(c.Request().Context())
	if err != nil {
		return nil, ErrInternalError
	}

	companies := make([]Company, len(companiesData))

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
