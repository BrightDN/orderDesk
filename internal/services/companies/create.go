package companies

import (
	"errors"
	"fmt"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

var ErrDuplicateName = errors.New("this name already exists")
var ErrDuplicateEmail = errors.New("this email already exists")

func (cs *CompanyService) Create(c echo.Context) (database.Company, error) {
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
				return database.Company{}, fmt.Errorf("%w", ErrDuplicateName)
			case "unique_email":
				return database.Company{}, fmt.Errorf("%w", ErrDuplicateEmail)
			}
		}
	}
	return comp, nil
}
