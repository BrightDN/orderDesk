package routing

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/pages"
	"github.com/labstack/echo/v4"
)

func (n *Navigation) appOrderHistory(c echo.Context) error {
	pageData := pages.PageData{
		Title:           "history",
		Type:            pages.BusinessType,
		SupplierDataURL: "app.new-order.get",
	}

	return c.Render(http.StatusOK, "/app/orderHistory", map[string]any{
		"pageData": pageData,
	})
}
