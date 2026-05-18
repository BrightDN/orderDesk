package invites

import (
	"fmt"
	"time"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/labstack/echo/v4"
)

func (is *InvitationService) Reactivate(c echo.Context) error {
	id, err := convertIDParam(c)
	if err != nil {
		return err
	}
	newExpiry := time.Now().Add(time.Hour * 48)
	if err := is.db.RenewInvite(c.Request().Context(), database.RenewInviteParams{
		ExpiresAt: newExpiry,
		ID:        id,
	}); err != nil {
		fmt.Printf("Error: %v", err)
		if flashErr := flash.Set(c, flash.Error, ErrInternalError.Error()); flashErr != nil {
			return flashErr
		}
		return err
	}
	return nil
}
