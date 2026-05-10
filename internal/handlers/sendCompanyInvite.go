package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/brightDN/orderDesk/internal/invites"
	"github.com/labstack/echo/v4"
)

var ErrCompanyCreationFailure = errors.New("Failed to create the company")

func (h *Handler) SendCompanyInvite(c echo.Context) error {
	org := h.App.Name
	orgmail := h.App.Cfg.MailAccount

	err := invites.SendCompany(h.App.Db, c, org, orgmail, h.App.Mailer)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCompanyCreationFailure, err)
	}

	return c.Redirect(http.StatusSeeOther, "/admin/companies/new")
}
