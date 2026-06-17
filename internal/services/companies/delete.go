package companies

import (
	"errors"

	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/labstack/echo/v4"
)

func (cs *CompanyService) Delete(c echo.Context, id int32) *errorHandling.AppError {
	err := cs.db.DeleteCompany(c.Request().Context(), id)
	if err != nil {
		return &errorHandling.AppError{
			Action:    "Deleting company",
			LogError:  err,
			UserError: errors.New("failed to delete company"),
		}
	}
	return nil
}
