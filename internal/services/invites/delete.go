package invites

import (
	"github.com/labstack/echo/v4"
)

func (is *InvitationService) Delete(c echo.Context, id int32) error {
	if err := is.db.DeleteInvite(c.Request().Context(), id); err != nil {
		return err
	}
	return nil
}
