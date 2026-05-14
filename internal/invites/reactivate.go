package invites

import (
	"time"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/labstack/echo/v4"
)

func Reactivate(db *database.Queries, c echo.Context, id int32) error {
	newExpiry := time.Now().Add(time.Hour * 48)
	if err := db.RenewInvite(c.Request().Context(), database.RenewInviteParams{
		ExpiresAt: newExpiry,
		ID:        id,
	}); err != nil {
		if flashErr := flash.Set(c, flash.Error, ErrInternalError.Error()); flashErr != nil {
			return flashErr
		}
		return err
	}
	return nil
}
