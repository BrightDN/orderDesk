package invites

import (
	"github.com/brightDN/orderDesk/internal/database"
	"github.com/labstack/echo/v4"
)

func Delete(db *database.Queries, c echo.Context, id int32) error {

	if err := db.DeleteInvite(c.Request().Context(), id); err != nil {
		return err
	}
	return nil
}
