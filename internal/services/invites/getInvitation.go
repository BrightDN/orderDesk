package invites

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/labstack/echo/v4"
)

func (is *InvitationService) GetInvitation(c echo.Context, token string) (Invitation, *errorHandling.AppError) {
	invite, err := is.db.GetInviteByToken(c.Request().Context(), token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Invitation{}, &errorHandling.AppError{
				Action:    "Validating invitation token",
				LogError:  fmt.Errorf("Invalid token provided: %s", token),
				UserError: errors.New("invalid invitation token"),
			}
		}
		return Invitation{}, &errorHandling.AppError{
			Action:    "Fetching invitation by token",
			LogError:  err,
			UserError: errors.New("failed to retrieve invitation"),
		}
	}
	if invite.ExpiresAt.Before(time.Now()) {
		return Invitation{}, &errorHandling.AppError{
			Action:    "Validating invitation expiry",
			LogError:  fmt.Errorf("Token expired at: %v", invite.ExpiresAt),
			UserError: errors.New("invitation token has expired"),
		}
	}
	if invite.UsedAt.Valid {
		return Invitation{}, &errorHandling.AppError{
			Action:    "Validating invitation usage",
			LogError:  fmt.Errorf("Token already used at: %v", invite.UsedAt),
			UserError: errors.New("invitation has already been used"),
		}
	}

	var invitation = Invitation{
		Type:    iType(invite.InviteType),
		Email:   invite.Email,
		Company: invite.CompanyName,
		Token:   invite.Token,
	}

	return invitation, nil
}
