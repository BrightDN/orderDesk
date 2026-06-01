package routing

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/pages"
	"github.com/labstack/echo/v4"
)

func (n *Navigation) appCompanySettings(c echo.Context) error {
	pageData := pages.PageData{
		Title: "company settings",
		Type:  pages.BusinessType,
	}
	return c.Render(http.StatusOK, "/app/companySettings", map[string]any{
		"pageData": pageData,
	})
}
