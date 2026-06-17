package invites

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/labstack/echo/v4"
)

func (inv *InvitationService) ValidateInvitation(c echo.Context, token, email string) (*database.Invite, *errorHandling.AppError) {
	invite, err := inv.db.ValidateInvite(c.Request().Context(), database.ValidateInviteParams{
		Token: token,
		Email: email,
	})

	if err == sql.ErrNoRows {
		return nil, &errorHandling.AppError{
			Action:    "Validating invitation credentials",
			LogError:  fmt.Errorf("Invalid invitation for token and email"),
			UserError: errors.New("invalid invitation"),
		}
	}
	if err != nil {
		return nil, &errorHandling.AppError{
			Action:    "Validating invitation",
			LogError:  err,
			UserError: errors.New("failed to validate invitation"),
		}
	}
	return &invite, nil
}
