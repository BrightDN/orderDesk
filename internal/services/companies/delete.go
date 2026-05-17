package companies

import (
	"github.com/brightDN/orderDesk/internal/database"
	"github.com/labstack/echo/v4"
)

func Delete(db *database.Queries, c echo.Context) error {
	id, err := convertIDParam(c)
	if err != nil {
		return err
	}

	return db.DeleteCompany(c.Request().Context(), id)
}
