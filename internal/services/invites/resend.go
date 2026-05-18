package invites

import (
	"fmt"
	"strings"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/brightDN/orderDesk/internal/services/mailer"
	"github.com/labstack/echo/v4"
)

func (is *InvitationService) Resend(c echo.Context, appname string) error {
	id, err := convertIDParam(c)
	if err != nil {
		return err
	}

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

	link := fmt.Sprintf("https://www.%s/invites/%s", strings.ToLower(appname), invite.Token)
	content := fmt.Sprintf("Hello,\n\nThank you for choosing %s.\nYour company account has been created successfully. To complete the setup process, please activate your account using the link below:\n%s\nPlease note that this activation link will expire in 48 hours.\nIf you did not request this account, you can safely ignore this email.\nBest regards,\nThe %s team",
		appname,
		link,
		appname)

	m := mailer.Mail{
		Receiver: invite.Email,
		Subject:  fmt.Sprintf("Activate your %s account", appname),
		Body:     content,
	}
	if err := is.mailService.Send(m); err != nil {
		return err
	}
	return nil
}
