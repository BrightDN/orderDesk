package routing

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/brightDN/orderDesk/internal/pages"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/brightDN/orderDesk/internal/shared/session"
	"github.com/labstack/echo/v4"
)

func (n *Navigation) appNewOrderDataPartial(c echo.Context) error {
	sname := c.Param("supplier-name")
	if strings.TrimSpace(sname) == "" {
		if logErr := errorHandling.Log_and_flash_trigger(c, errorHandling.AppError{
			Action:    "Retrieving parameter for supplier-name",
			LogError:  fmt.Errorf("Param field: \"Suppliername\" was empty"),
			UserError: fmt.Errorf("error: Failed to retrieve suppliername"),
		}); logErr != nil {
			return logErr
		}
		return c.NoContent(http.StatusNoContent)
	}

	compID, ok, err := session.GetValue[int32](c, session.CompanyIDKey)
	if err != nil || !ok {
		if logErr := errorHandling.Log_and_flash(c, errorHandling.AppError{
			Action:    "Retrieving companyID value from session",
			LogError:  err,
			UserError: fmt.Errorf("Failed to read company ID"),
		}); logErr != nil {
			return logErr
		}
		return c.Redirect(http.StatusSeeOther, Logout)
	}

	supplier, aerr := n.app.Services.Suppliers.GetSupplierByNameAndCompanyID(c, sname, compID)
	if aerr != nil {
		if logErr := errorHandling.Log_and_flash_trigger(c, *aerr); logErr != nil {
			return logErr
		}
		return c.NoContent(http.StatusNoContent)
	}

	products, aerr := n.app.Services.Suppliers.GetProducts(c, supplier.ID)
	if aerr != nil {
		if logErr := errorHandling.Log_and_flash_trigger(c, *aerr); logErr != nil {
			return logErr
		}
		return c.NoContent(http.StatusNoContent)
	}

	pageData := pages.PageData{
		Title:           "new order",
		Type:            pages.BusinessType,
		SupplierDataURL: "app.new-order.get",
	}

	fmt.Println(c.Request().RequestURI)
	return c.Render(http.StatusOK, "partials/itemList", map[string]any{
		"products": products,
		"supplier": supplier,
		"pageData": pageData,
	})
}
