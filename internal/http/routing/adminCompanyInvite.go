package routing

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/pages"
	"github.com/labstack/echo/v4"
)

func (n *Navigation) adminCompanyInvite(c echo.Context) error {
	invs := n.app.Services.Invitations.GetCompanyInvites(c)
	pageData := pages.PageData{
		Title: "Invite companies",
		Type:  pages.OwnerType,
	}
	return c.Render(http.StatusOK, "adminCompanyInvites", map[string]any{
		"invites":  invs,
		"pageData": pageData,
		"csrf":     c.Get("csrf"),
	})
}
