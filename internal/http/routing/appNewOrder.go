package routing

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/pages"
	"github.com/labstack/echo/v4"
)

func (n *Navigation) appNewOrder(c echo.Context) error {
	pageData := pages.PageData{
		Title: "new order",
		Type:  pages.BusinessType,
	}

	return c.Render(http.StatusOK, "app/newOrder", map[string]any{
		"pageData": pageData,
	})
}
