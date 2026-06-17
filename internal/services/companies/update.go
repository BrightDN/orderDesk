package companies

import (
	"errors"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/labstack/echo/v4"
)

func (cs *CompanyService) Update(c echo.Context, id int32) *errorHandling.AppError {
	name := c.Request().PostFormValue("name")
	email := c.Request().PostFormValue("email")

	if err := cs.db.UpdateCompany(c.Request().Context(), database.UpdateCompanyParams{
		ID:    id,
		Name:  name,
		Email: email,
	}); err != nil {
		return &errorHandling.AppError{
			Action:    "Updating company information",
			LogError:  err,
			UserError: errors.New("failed to update company"),
		}
	}

	if c.Request().PostFormValue("status") == "inactive" {
		if err := cs.db.DeleteCompany(c.Request().Context(), id); err != nil {
			return &errorHandling.AppError{
				Action:    "Deactivating company",
				LogError:  err,
				UserError: errors.New("failed to deactivate company"),
			}
		}
	}

	return nil
}
