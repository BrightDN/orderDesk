package routing

import (
	"errors"
	"net/http"
	"strings"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/brightDN/orderDesk/internal/pages"
	"github.com/brightDN/orderDesk/internal/shared/logging"
	"github.com/brightDN/orderDesk/internal/shared/session"
	"github.com/labstack/echo/v4"
)

const ACTION_GETSUPPLIER = "Retrieving supplier after supplierlist click"

var ErrFailedToRetrieveData = errors.New("error: failed to retrieve data")

func (n *Navigation) appGetSupplier(c echo.Context) error {
	sname := c.Param("supplier-name")
	if strings.TrimSpace(sname) == "" {
		logging.ErrorLog(ACTION_GETSUPPLIER, "Paramfield was empty")
		if flashErr := flash.Set(c, flash.Error, "error: Failed to retrieve suppliername"); flashErr != nil {
			return flashErr
		}
		return c.NoContent(http.StatusNoContent)
	}

	compID, ok, err := session.GetValue[int32](c, session.CompanyIDKey)
	if err != nil || !ok {
		logging.ErrorLog(ACTION_GETSUPPLIER, err.Error())
		if flashErr := flash.Set(c, flash.Error, err.Error()); flashErr != nil {
			return flashErr
		}
		return c.Redirect(http.StatusSeeOther, Login)
	}

	supplier, aerr := n.app.Services.Suppliers.GetSupplierByNameAndCompanyID(c, sname, compID)
	if aerr != nil {
		if flashErr := flash.Set(c, flash.Error, ErrFailedToRetrieveData.Error()); flashErr != nil {
			return flashErr
		}
		return c.NoContent(http.StatusNoContent)
	}
	products, aerr := n.app.Services.Suppliers.GetProducts(c, supplier.ID)
	if aerr != nil {
		logging.ErrorLog(ACTION_GETSUPPLIER, aerr.Error())
		if flashErr := flash.Set(c, flash.Error, ErrFailedToRetrieveData.Error()); flashErr != nil {
			return flashErr
		}
		return c.NoContent(http.StatusNoContent)
	}

	pageData := pages.PageData{
		Title: "suppliers",
		Type:  pages.BusinessType,
	}
	return c.Render(http.StatusOK, "partials/fullSupplierData", map[string]any{
		"supplier": supplier,
		"products": products,
		"pageData": pageData,
	})
}
