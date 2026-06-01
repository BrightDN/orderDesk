package routing

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/pages"
	"github.com/labstack/echo/v4"
)

func (n *Navigation) supportContact(c echo.Context) error {
	pageData := pages.PageData{
		Title: "Contact support",
		Type:  pages.BusinessType,
	}
	return c.Render(http.StatusOK, "/support/contact", map[string]any{
		"pageData": pageData,
	})
}
