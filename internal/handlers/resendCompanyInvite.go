package handlers

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/brightDN/orderDesk/internal/invites"
	"github.com/labstack/echo/v4"
)

func (h *Handler) ResendCompanyInvite(c echo.Context) error {
	if err := invites.Resend(h.App.Db, c, h.App.Mailer, h.App.Name, h.App.Cfg.MailAccount); err != nil {
		if flashErr := flash.Trigger(c, flash.Error, err.Error()); flashErr != nil {
			return flashErr
		}
		return c.NoContent(http.StatusNoContent)
	}

	if err := flash.Trigger(c, flash.Pass, "Company invite email resent."); err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
