package handlers

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/brightDN/orderDesk/internal/services/companies/orders"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/brightDN/orderDesk/internal/shared/session"
	"github.com/labstack/echo/v4"
)

func (h *Handler) sendOrder(c echo.Context) error {

	var order orders.OrderResponse

	if err := c.Bind(&order); err != nil {
		return returnFeedback(c, flash.Error, nil, "error: Malformed request")
	}

	compID, ok, err := session.GetValue[int32](c, session.CompanyIDKey)
	if err != nil || !ok {

		return returnFeedback(c, flash.Error, err, "")
	}
	userID, ok, err := session.GetValue[int32](c, session.UserIDKey)
	if err != nil || !ok {
		return returnFeedback(c, flash.Error, err, "")
	}

	supplier, err := h.App.Services.Suppliers.GetSupplierByNameAndCompanyID(c, order.SupplierName, compID)
	if err != nil {
		return returnFeedback(c, flash.Error, err, "")
	}

	employee, err := h.App.Services.Companies.GetEmployee(c, userID)
	if err != nil {
		return returnFeedback(c, flash.Error, err, "")
	}

	orderData, err := h.App.Services.Orders.GetOrderData(order, supplier, employee)
	if err != nil {
		return returnFeedback(c, flash.Error, err, "")
	}

	mailData, err := h.App.Services.Orders.GetOrderMailData(c, supplier.ID, *orderData)
	if err != nil {
		return returnFeedback(c, flash.Error, err, "")
	}

	if err := h.App.Services.Mailer.SendOrder(mailData); err != nil {
		return returnFeedback(c, flash.Error, err, "")
	}
	orderData = orderData
	mailData = mailData

	return returnFeedback(c, flash.Pass, nil, "Order has been sent")
}

func returnFeedback(c echo.Context, t flash.MessageType, err *errorHandling.AppError, msg string) error {
	m := msg
	if err != nil {
		if logErr := errorHandling.Log_and_flash(c, *err); logErr != nil {
			return logErr
		}
		m = err.UserError.Error()
	}
	feedback := flash.Flash{
		Type:    t,
		Message: m,
	}
	return c.Render(http.StatusOK, "/components/feedback", map[string]any{
		"feedback": feedback,
	})
}
