package invites

import (
	"errors"
	"fmt"
	"time"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/labstack/echo/v4"
)

func (is *InvitationService) Reactivate(c echo.Context, id int32) *errorHandling.AppError {
	newExpiry := time.Now().Add(time.Hour * 48)
	if err := is.db.RenewInvite(c.Request().Context(), database.RenewInviteParams{
		ExpiresAt: newExpiry,
		ID:        id,
	}); err != nil {
		return &errorHandling.AppError{
			Action:    "Reactivating expired invitation",
			LogError:  fmt.Errorf("Failed to renew invitation %d: %v", id, err),
			UserError: errors.New("failed to reactivate invitation"),
		}
	}
	return nil
}
