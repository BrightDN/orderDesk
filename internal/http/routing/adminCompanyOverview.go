package routing

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/pages"
	"github.com/brightDN/orderDesk/internal/services/companies"
	"github.com/labstack/echo/v4"
)

func (n *Navigation) adminCompanyOverview(c echo.Context) error {
	companies, err := companies.GetCompanies(n.db, c)
	if err != nil {
		return err
	}

	pageData := pages.PageData{
		Title: "Companies",
		Type:  pages.OwnerType,
	}
	return c.Render(http.StatusOK, "adminCompanyOverview", map[string]any{
		"pageData":  pageData,
		"companies": companies,
	})
}
