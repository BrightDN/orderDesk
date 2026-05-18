package routing

import (
	"errors"
	"net/http"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/brightDN/orderDesk/internal/pages"
	"github.com/labstack/echo/v4"
)

var ErrCompanyDetailsLoading = errors.New("the requested company could not be loaded")

func (n *Navigation) adminCompanyDetails(c echo.Context) error {
	company, err := n.app.Services.Companies.GetCompany(c)
	if err != nil {
		if flashErr := flash.Set(c, flash.Error, ErrCompanyDetailsLoading.Error()); flashErr != nil {
			return flashErr
		}
		return c.Redirect(http.StatusSeeOther, "/admin/companies/overview")
	}

	pageData := pages.PageData{
		Title: "Company details",
		Type:  pages.OwnerType,
	}
	return c.Render(http.StatusOK, "adminCompanyDetails", map[string]any{
		"pageData": pageData,
		"company":  company,
	})
}
