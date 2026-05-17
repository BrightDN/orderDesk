package handlers

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/brightDN/orderDesk/internal/services/invites"
	"github.com/labstack/echo/v4"
)

const adminNewCompanyPath = "/admin/companies/invites"

func (h *Handler) sendCompanyInvite(c echo.Context) error {
	err := invites.SendCompany(h.App.Db, c, h.App.Name, h.App.Cfg.Mail.Email, h.App.Mailer)
	if err != nil {
		if flashErr := flash.Set(c, flash.Error, err.Error()); flashErr != nil {
			return flashErr
		}
		return c.Redirect(http.StatusSeeOther, adminNewCompanyPath)
	}

	if err := flash.Set(c, flash.Pass, "Company invite created and sent."); err != nil {
		return err
	}
	return c.Redirect(http.StatusSeeOther, adminNewCompanyPath)
}
