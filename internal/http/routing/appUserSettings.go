package routing

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/pages"
	"github.com/labstack/echo/v4"
)

func (n *Navigation) appUserSettings(c echo.Context) error {
	pageData := pages.PageData{
		Title: "user settings",
		Type:  pages.BusinessType,
	}
	return c.Render(http.StatusOK, "app/userSettings", map[string]any{
		"pageData": pageData,
	})
}
