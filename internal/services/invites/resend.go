package invites

import (
	"fmt"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/brightDN/orderDesk/internal/services/mailer"
	"github.com/labstack/echo/v4"
)

func (is *InvitationService) Resend(c echo.Context, id int32) error {
	invite, err := is.db.GetInvite(c.Request().Context(), id)
	if err != nil {
		if flashErr := flash.Set(c, flash.Error, ErrInternalError.Error()); flashErr != nil {
			return flashErr
		}
		return err
	}

	if invite.UsedAt.Valid {
		if flashErr := flash.Set(c, flash.Error, ErrAlreadyAccepted.Error()); flashErr != nil {
			return flashErr
		}
		return err
	}

	content := is.getCompanyInvMail(invite.Token)

	m := mailer.Mail{
		Receiver: invite.Email,
		Subject:  fmt.Sprintf("Activate your %s account", is.identity.AppName),
		Body:     content,
	}
	if err := is.mailService.Send(m); err != nil {
		return err
	}
	return nil
}
