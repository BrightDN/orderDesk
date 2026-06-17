package companies

import (
	"errors"
	"fmt"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func (cs *CompanyService) Create(c echo.Context) (database.Company, *errorHandling.AppError) {
	mail := c.Request().PostFormValue("email")
	name := c.Request().PostFormValue("company-name")
	comp, err := cs.db.CreateCompany(c.Request().Context(), database.CreateCompanyParams{
		Name:  name,
		Email: mail,
	})

	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) &&
			pqErr.Code == "23505" {
			switch pqErr.Constraint {
			case "unique_name":
				return database.Company{}, &errorHandling.AppError{
					Action:    "Creating company",
					LogError:  fmt.Errorf("Duplicate company name: %s", name),
					UserError: errors.New("this name already exists"),
				}
			case "unique_email":
				return database.Company{}, &errorHandling.AppError{
					Action:    "Creating company",
					LogError:  fmt.Errorf("Duplicate company email: %s", mail),
					UserError: errors.New("this email already exists"),
				}
			}
		}
		return database.Company{}, &errorHandling.AppError{
			Action:    "Creating company",
			LogError:  err,
			UserError: errors.New("failed to create company"),
		}
	}
	return comp, nil
}
