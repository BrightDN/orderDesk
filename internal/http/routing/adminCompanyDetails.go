package routing

import (
	"errors"
	"net/http"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/brightDN/orderDesk/internal/pages"
	"github.com/brightDN/orderDesk/internal/shared/parse"
	"github.com/labstack/echo/v4"
)

var ErrCompanyDetailsLoading = errors.New("the requested company could not be loaded")

func (n *Navigation) adminCompanyDetails(c echo.Context) error {
	id, err := parse.Int32(c.Param("id"))
	company, err := n.app.Services.Companies.GetCompany(c, id)
	if err != nil {
		if flashErr := flash.Set(c, flash.Error, ErrCompanyDetailsLoading.Error()); flashErr != nil {
			return flashErr
		}
		return c.Redirect(http.StatusSeeOther, "/admin/companies/overview")
	}

	employees, err := n.app.Services.Companies.GetCompanyEmployees(c, id)

	pageData := pages.PageData{
		Title: "Company details",
		Type:  pages.OwnerType,
	}
	return c.Render(http.StatusOK, "adminCompanyDetails", map[string]any{
		"pageData":  pageData,
		"company":   company,
		"employees": employees,
	})
}
