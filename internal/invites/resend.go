package invites

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/brightDN/orderDesk/internal/mailer"
	"github.com/labstack/echo/v4"
	"github.com/wneessen/go-mail"
)

func Resend(db *database.Queries, c echo.Context, mailc *mail.Client, appname, appmail string) error {
	id := c.Param("id")
	if len(strings.TrimSpace(id)) == 0 {
		if flashErr := flash.Set(c, flash.Error, ErrUnexpectedValue.Error()); flashErr != nil {
			return flashErr
		}
	}

	nid, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		if flashErr := flash.Set(c, flash.Error, ErrUnexpectedValue.Error()); flashErr != nil {
			return flashErr
		}
		return err
	}

	invite, err := db.GetInvite(c.Request().Context(), int32(nid))
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
		Sender:   appmail,
		Subject:  fmt.Sprintf("Activate your %s account", appname),
		Body:     content,
	}
	if err := mailer.SendMail(m, mailc); err != nil {
		return err
	}
	return nil
}
