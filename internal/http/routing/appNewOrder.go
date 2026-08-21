package routing

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/pages"
	"github.com/brightDN/orderDesk/internal/services/companies/suppliers"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/brightDN/orderDesk/internal/shared/session"
	"github.com/labstack/echo/v4"
)

func (n *Navigation) appNewOrder(c echo.Context) error {
	pageData := pages.PageData{
		Title:           "new order",
		Type:            pages.BusinessType,
		SupplierDataURL: "app.new-order.get",
	}

	id, ok, err := session.GetValue[int32](c, session.CompanyIDKey)
	if err != nil {
		if logErr := errorHandling.Log_and_flash(c, *err); logErr != nil {
			return logErr
		}
		return c.Redirect(http.StatusSeeOther, Logout)
	}
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "company not found")
	}

	suppl, err := n.app.Services.Suppliers.GetAllByCompany(c, id)
	if err != nil {
		if logErr := errorHandling.Log_and_flash(c, *err); logErr != nil {
			return logErr
		}
		return c.Redirect(http.StatusSeeOther, Logout)
	}
	var products []suppliers.Products
	var supp suppliers.Supplier
	if len(suppl) > 0 {
		supp = suppl[0]
		products, err = n.app.Services.Suppliers.GetProducts(c, suppl[0].ID)
		if err != nil {
			if err := errorHandling.Log_and_flash(c, *err); err != nil {
				return err
			}
		}
	}

	return c.Render(http.StatusOK, NeworderPage, map[string]any{
		"pageData":    pageData,
		"supplier":    supp,
		"products":    products,
		"permissions": appPermissions(c),
	})
}
