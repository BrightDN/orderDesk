package invites

import (
	"errors"
	"fmt"

	"github.com/brightDN/orderDesk/internal/services/mailer"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/labstack/echo/v4"
)

func (is *InvitationService) Resend(c echo.Context, id int32) *errorHandling.AppError {
	invite, err := is.db.GetInvite(c.Request().Context(), id)
	if err != nil {
		return &errorHandling.AppError{
			Action:    "Fetching invitation for resend",
			LogError:  fmt.Errorf("Failed to fetch invitation %d: %v", id, err),
			UserError: errors.New("failed to resend invitation"),
		}
	}

	if invite.UsedAt.Valid {
		return &errorHandling.AppError{
			Action:    "Resending invitation",
			LogError:  fmt.Errorf("Invitation %d already accepted", id),
			UserError: errors.New("invitation has already been accepted"),
		}
	}

	content := is.getCompanyInvMail(invite.Token)

	m := mailer.Mail{
		Receiver: invite.Email,
		Subject:  fmt.Sprintf("Activate your %s account", is.identity.AppName),
		Body:     content,
	}
	if err := is.mailService.Send(m); err != nil {
		return &errorHandling.AppError{
			Action:    "Sending invitation email",
			LogError:  err,
			UserError: errors.New("failed to send invitation email"),
		}
	}
	return nil
}
