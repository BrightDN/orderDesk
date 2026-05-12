package handlers

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/invites"
	"github.com/labstack/echo/v4"
)

func (h *Handler) NavAdminNewCompany(c echo.Context) error {
	invs := invites.GetCompanyInvites(h.App.Db, c, h.App.Name)
	return c.Render(http.StatusOK, "inviteCompany", map[string]any{
		"invites": invs,
	})
}
