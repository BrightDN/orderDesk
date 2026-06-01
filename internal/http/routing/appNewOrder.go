package routing

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/brightDN/orderDesk/internal/pages"
	"github.com/brightDN/orderDesk/internal/shared/session"
	"github.com/labstack/echo/v4"
)

func (n *Navigation) appNewOrder(c echo.Context) error {
	pageData := pages.PageData{
		Title: "new order",
		Type:  pages.BusinessType,
	}

	empl := c.Get("employee")
	if empl == nil {
		flashErr := flash.Set(c, flash.Error, "Failed to load employee data. Please log in again.")
		if flashErr != nil {
			return flashErr
		}
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	multiCompany, _, err := session.GetValue[bool](c, session.MultiCompKey)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "app/newOrder", map[string]any{
		"pageData":     pageData,
		"employee":     empl,
		"multiCompany": multiCompany,
	})
}
