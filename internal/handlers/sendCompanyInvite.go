package handlers

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/brightDN/orderDesk/internal/invites"
	"github.com/labstack/echo/v4"
)

const adminNewCompanyPath = "/admin/companies/new"

func (h *Handler) SendCompanyInvite(c echo.Context) error {
	err := invites.SendCompany(h.App.Db, c, h.App.Name, h.App.Cfg.MailAccount, h.App.Mailer)
	if err != nil {
		if flashErr := flash.Set(c, "error", err.Error()); flashErr != nil {
			return flashErr
		}
		return c.Redirect(http.StatusSeeOther, adminNewCompanyPath)
	}

	if err := flash.Set(c, "pass", "Company invite created and sent."); err != nil {
		return err
	}
	return c.Redirect(http.StatusSeeOther, adminNewCompanyPath)
}
