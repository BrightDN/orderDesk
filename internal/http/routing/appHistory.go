package routing

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/pages"
	"github.com/brightDN/orderDesk/internal/services/companies/orders"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/brightDN/orderDesk/internal/shared/session"
	"github.com/labstack/echo/v4"
)

func (n *Navigation) appOrderHistory(c echo.Context) error {
	pageData := pages.PageData{
		Title:           "history",
		Type:            pages.BusinessType,
		SupplierDataURL: "app.history.get",
	}

	compID, ok, err := session.GetValue[int32](c, session.CompanyIDKey)
	if err != nil {
		if logErr := errorHandling.Log_and_flash(c, *err); logErr != nil {
			return logErr
		}
		return c.Redirect(http.StatusSeeOther, Logout)
	}
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "company not found")
	}

	suppl, err := n.app.Services.Suppliers.GetAllByCompany(c, compID)
	if err != nil {
		if logErr := errorHandling.Log_and_flash(c, *err); logErr != nil {
			return logErr
		}
		return c.Redirect(http.StatusSeeOther, Logout)
	}

	var orderHistory []orders.OrderHistoryLog
	if len(suppl) > 0 {
		orderHistory, err = n.app.Services.Orders.GetOrderHistoryForSupplier(c, suppl[0].ID, compID)
	}

	return c.Render(http.StatusOK, "/app/orderHistory", map[string]any{
		"pageData":     pageData,
		"orderHistory": orderHistory,
		"supplierName": suppl[0].Name,
		"permissions":  appPermissions(c),
	})
}
