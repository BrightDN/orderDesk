package companies

import (
	"github.com/labstack/echo/v4"
)

func (cs *CompanyService) Delete(c echo.Context) error {
	id, err := convertIDParam(c)
	if err != nil {
		return err
	}

	return cs.db.DeleteCompany(c.Request().Context(), id)
}
