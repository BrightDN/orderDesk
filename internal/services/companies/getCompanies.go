package companies

import (
	"errors"

	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/labstack/echo/v4"
)

func (cs *CompanyService) GetCompanies(c echo.Context) ([]Company, *errorHandling.AppError) {
	companiesData, err := cs.db.GetCompanies(c.Request().Context())
	if err != nil {
		return nil, &errorHandling.AppError{
			Action:    "Fetching companies",
			LogError:  err,
			UserError: errors.New("failed to fetch companies"),
		}
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
