package invites

import (
	"errors"
	"fmt"

	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/labstack/echo/v4"
)

func (is *InvitationService) Delete(c echo.Context, id int32) *errorHandling.AppError {
	if err := is.db.DeleteInvite(c.Request().Context(), id); err != nil {
		return &errorHandling.AppError{
			Action:    "Deleting invitation",
			LogError:  fmt.Errorf("Failed to delete invitation %d: %v", id, err),
			UserError: errors.New("failed to delete invitation"),
		}
	}
	return nil
}
