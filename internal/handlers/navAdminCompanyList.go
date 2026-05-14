package handlers

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/companies"
	"github.com/brightDN/orderDesk/internal/pages"
	"github.com/labstack/echo/v4"
)

func (h *Handler) NavAdminCompanyList(c echo.Context) error {
	companies := companies.GetCompanies(h.App.Db, c)

	pageData := pages.PageData{
		Title: "Companies",
		Type:  pages.OwnerType,
	}
	return c.Render(http.StatusOK, "SACompanyPanel", map[string]any{
		"pageData":  pageData,
		"companies": companies,
	})
}
