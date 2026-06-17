package companies

import (
	"errors"
	"fmt"

	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/labstack/echo/v4"
)

func (cs *CompanyService) GetCompany(c echo.Context, id int32) (Company, *errorHandling.AppError) {
	data, err := cs.db.GetCompany(c.Request().Context(), id)
	if err != nil {
		return Company{}, &errorHandling.AppError{
			Action:    "Fetching company details",
			LogError:  fmt.Errorf("Company not found with ID %d: %v", id, err),
			UserError: errors.New("failed to fetch company"),
		}
	}

	company := Company{
		ID:        data.ID,
		Name:      data.Name,
		Email:     data.Email,
		CreatedAt: data.CreatedAt.Format("02-01-2006"),
		IsDeleted: data.DeletedAt.Valid,
	}

	return company, nil
}
