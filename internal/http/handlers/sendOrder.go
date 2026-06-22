package handlers

import (
	"fmt"
	"net/http"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/labstack/echo/v4"
)

func (h *Handler) sendOrder(c echo.Context) error {

	var order Order

	if err := c.Bind(&order); err != nil {
		return err
	}

	fmt.Println(order.SupplierID)

	for _, item := range order.Items {
		fmt.Println(item.ProductName, item.Qty)
	}

	return returnFeedback(c, flash.Pass, "Order has been sent")
}

type Order struct {
	SupplierID int         `json:"supplierID"`
	Items      []OrderItem `json:"items"`
}

type OrderItem struct {
	ProductName string `json:"productName"`
	Qty         int    `json:"qty"`
}

func returnFeedback(c echo.Context, t flash.MessageType, msg string) error {
	feedback := flash.Flash{
		Type:    t,
		Message: msg,
	}
	return c.Render(http.StatusOK, "/components/feedback", map[string]any{
		"feedback": feedback,
	})
}
