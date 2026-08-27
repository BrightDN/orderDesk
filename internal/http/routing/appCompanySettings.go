package routing

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/pages"
	"github.com/brightDN/orderDesk/internal/services/companies"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/brightDN/orderDesk/internal/shared/session"
	"github.com/labstack/echo/v4"
)

func (n *Navigation) appCompanySettings(c echo.Context) error {
	pageData := pages.PageData{
		Title: "company settings",
		Type:  pages.BusinessType,
	}

	id, ok, aerr := session.GetValue[int32](c, session.CompanyIDKey)
	if !ok || aerr != nil {
		if err := errorHandling.Log_and_flash(c, *aerr); err != nil {
			return err
		}
		return c.Redirect(http.StatusOK, Logout)
	}
	employees, aerr := n.app.Services.Companies.GetCompanyEmployees(c, id)
	if aerr != nil {
		if err := errorHandling.Log_and_flash(c, *aerr); err != nil {
			return err
		}
		return c.Redirect(http.StatusOK, Logout)
	}
	currentEmployee := c.Get("employee").(companies.Employee)

	return c.Render(http.StatusOK, "app/companySettings", map[string]any{
		"pageData":    pageData,
		"permissions": appPermissions(c),
		"employee":    currentEmployee,
		"employees":   employees,
	})
}
