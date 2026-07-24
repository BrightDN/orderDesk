package orders

import (
	"fmt"
	"strings"

	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/labstack/echo/v4"
)

func (os *OrderService) GetOrderMailData(c echo.Context, supplierID int32, orderData Order) (*OrderMailData, *errorHandling.AppError) {
	data, err := os.queries.GetOrderMailData(c.Request().Context(), supplierID)
	if err != nil {
		return nil, &errorHandling.AppError{
			Action:    "Retrieving order mail from database",
			LogError:  fmt.Errorf("error occured while retrieving mail information from the database: %v", err),
			UserError: fmt.Errorf("error: something went wrong and we could not complete the request. Please try again or contact support"),
		}
	}

	body := os.replaceUserVariable(data.MailContent, orderData.Sender.Contact)
	body = os.replaceCompanyVariable(body, orderData.Sender.Org)

	mailData := OrderMailData{
		Subject: data.Subject,
		Body:    body,
		Order:   &orderData,
	}

	return &mailData, nil
}

type OrderMailData struct {
	Subject string
	Body    string
	Order   *Order
}

func (os *OrderService) replaceUserVariable(target, replacement string) string {
	return strings.Replace(target, "{{ user }}", replacement, -1)
}

func (os *OrderService) replaceCompanyVariable(target, replacement string) string {
	return strings.Replace(target, "{{ company }}", replacement, -1)
}
