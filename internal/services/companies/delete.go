package companies

import (
	"github.com/labstack/echo/v4"
)

func (cs *CompanyService) Delete(c echo.Context, id int32) error {
	return cs.db.DeleteCompany(c.Request().Context(), id)
}
