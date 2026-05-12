package handlers

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/invites"
	"github.com/brightDN/orderDesk/internal/pages"
	"github.com/labstack/echo/v4"
)

func (h *Handler) NavAdminNewCompany(c echo.Context) error {
	invs := invites.GetCompanyInvites(h.App.Db, c, h.App.Name)
	pageData := pages.PageData{
		Title: "Invite companies",
		Type:  pages.OwnerType,
	}
	return c.Render(http.StatusOK, "inviteCompany", map[string]any{
		"invites":  invs,
		"pageData": pageData,
	})
}
