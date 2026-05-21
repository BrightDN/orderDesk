package companies

import (
	"github.com/brightDN/orderDesk/internal/database"
	"github.com/labstack/echo/v4"
)

func (cs *CompanyService) Update(c echo.Context, id int32) error {
	name := c.Request().PostFormValue("name")
	email := c.Request().PostFormValue("email")

	if err := cs.db.UpdateCompany(c.Request().Context(), database.UpdateCompanyParams{
		ID:    id,
		Name:  name,
		Email: email,
	}); err != nil {
		return ErrInternalError
	}

	if c.Request().PostFormValue("status") == "inactive" {
		if err := cs.db.DeleteCompany(c.Request().Context(), id); err != nil {
			return ErrInternalError
		}
	}

	return nil
}
