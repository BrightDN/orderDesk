package routing

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/brightDN/orderDesk/internal/shared/session"
	"github.com/labstack/echo/v4"
)

const (
	htmxRequestHeader  = "HX-Request"
	htmxRedirectHeader = "HX-Redirect"
)

func (n *Navigation) appOrderDownload(c echo.Context) error {
	orderID := strings.TrimSpace(c.Param("order-id"))
	if orderID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "order ID is required")
	}

	// An XHR cannot invoke the browser's native attachment download. Tell HTMX
	// to navigate to this same protected route, which makes the next request a
	// normal browser navigation and therefore downloads the PDF.
	if c.Request().Header.Get(htmxRequestHeader) == "true" {
		c.Response().Header().Set(htmxRedirectHeader, c.Request().URL.Path)
		return c.NoContent(http.StatusNoContent)
	}

	companyID, ok, aerr := session.GetValue[int32](c, session.CompanyIDKey)
	if aerr != nil || !ok {
		if aerr == nil {
			aerr = &errorHandling.AppError{
				Action:    "Retrieving company ID for order download",
				LogError:  fmt.Errorf("company ID missing from session"),
				UserError: fmt.Errorf("error: failed to retrieve company information"),
			}
		}
		if logErr := errorHandling.Log_and_flash(c, *aerr); logErr != nil {
			return logErr
		}
		return c.Redirect(http.StatusSeeOther, Logout)
	}

	order, aerr := n.app.Services.Orders.GenerateOrderPDFForDownload(c, orderID, companyID)
	if aerr != nil {
		return echo.NewHTTPError(http.StatusNotFound, aerr.UserError.Error())
	}

	c.Response().Header().Set(echo.HeaderContentType, "application/pdf")
	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename="+strconv.Quote(order.ID+".pdf"))
	return c.Stream(http.StatusOK, "application/pdf", order.PDF)
}
