package routing

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/pages"
	"github.com/labstack/echo/v4"
)

func (n *Navigation) authSelectCompany(c echo.Context) error {
	pageData := pages.PageData{
		Title: "Select company",
		Type:  pages.BusinessType,
	}
	return c.Render(http.StatusOK, "auth/select-company", map[string]any{
		"pageData": pageData,
	})
}
