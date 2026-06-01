package routing

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/pages"
	"github.com/labstack/echo/v4"
)

func (n *Navigation) appSuppliers(c echo.Context) error {
	pageData := pages.PageData{
		Title: "suppliers",
		Type:  pages.BusinessType,
	}

	return c.Render(http.StatusOK, "app/suppliers", map[string]any{
		"pageData": pageData,
	})
}
