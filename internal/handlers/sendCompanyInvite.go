package handlers

import (
	"fmt"
	"net/http"

	"github.com/brightDN/orderDesk/internal/mailer"
	"github.com/labstack/echo/v4"
)

func (h *Handler) SendCompanyInvite(c echo.Context) error {
	org := h.App.Name
	subject := org + " account activation for " + c.Request().PostFormValue("company-name")
	sm := h.App.Cfg.MailAccount
	rm := c.Request().PostFormValue("email")
	link := "#"
	content := fmt.Sprintf("Hello\n\nThank you for choosing %s.\nYou can activate your company account through the link below.\n%s\n\nThis link will expire in 2 days.\n\nSincerely,\nThe %s team",
		org,
		link,
		org)

	if err := mailer.SendMail(rm, sm, subject, content, h.App.Mailer); err != nil {
		return fmt.Errorf("Could not send the mail: %v", err)
	}
	return c.Redirect(http.StatusSeeOther, "/admin/controlpanel")
}
