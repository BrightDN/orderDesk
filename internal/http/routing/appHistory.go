package routing

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/pages"
	"github.com/labstack/echo/v4"
)

func (n *Navigation) appOrderHistory(c echo.Context) error {
	pageData := pages.PageData{
		Title: "new order",
		Type:  pages.BusinessType,
	}
	return c.Render(http.StatusOK, "/app/orderHistory", map[string]any{
		"pageData": pageData,
	})
}
