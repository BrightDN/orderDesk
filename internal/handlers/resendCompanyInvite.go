package handlers

import (
	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/brightDN/orderDesk/internal/invites"
	"github.com/labstack/echo/v4"
)

func (h *Handler) ResendCompanyInvite(c echo.Context) error {
	if err := invites.Resend(h.App.Db, c, h.App.Mailer, h.App.Name, h.App.Cfg.MailAccount); err != nil {
		if flashErr := flash.Set(c, "error", err.Error()); flashErr != nil {
			return flashErr
		}
		return err
	}
	return nil
}
