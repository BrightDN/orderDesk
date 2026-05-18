package invites

import (
	"github.com/labstack/echo/v4"
)

func (is *InvitationService) Delete(c echo.Context) error {
	id, err := convertIDParam(c)
	if err != nil {
		return err
	}

	if err := is.db.DeleteInvite(c.Request().Context(), id); err != nil {
		return err
	}
	return nil
}
