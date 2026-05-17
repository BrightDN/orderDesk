package routing

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/pages"
	"github.com/brightDN/orderDesk/internal/services/invites"
	"github.com/labstack/echo/v4"
)

func (n *Navigation) adminCompanyInvite(c echo.Context) error {
	invs := invites.GetCompanyInvites(n.db, c)
	pageData := pages.PageData{
		Title: "Invite companies",
		Type:  pages.OwnerType,
	}
	return c.Render(http.StatusOK, "adminCompanyInvites", map[string]any{
		"invites":  invs,
		"pageData": pageData,
	})
}
