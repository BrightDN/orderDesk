package routing

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/brightDN/orderDesk/internal/shared/session"
	"github.com/labstack/echo/v4"
)

func (n *Navigation) appHistoryDataPartial(c echo.Context) error {
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

	orderHistory, err := n.app.Services.Orders.GetOrderHistoryForSupplier(c, supplier.ID, compID)
	if err != nil {
		if logErr := errorHandling.Log_and_flash(c, *err); logErr != nil {
			return logErr
		}
	}

	return c.Render(http.StatusOK, "components/historyTable", map[string]any{
		"orderHistory": orderHistory,
		"isPartial":    true,
		"supplierName": supplier.Name,
		"permissions":  appPermissions(c),
	})
}
