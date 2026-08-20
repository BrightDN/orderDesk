package orders

import (
	"errors"
	"fmt"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/labstack/echo/v4"
)

// GenerateOrderPDFForDownload rebuilds an order from its immutable order-item
// records, so a historical download matches the order that was placed.
func (os *OrderService) GenerateOrderPDFForDownload(c echo.Context, orderID string, companyID int32) (*Order, *errorHandling.AppError) {
	rows, err := os.queries.GetOrderForDownload(c.Request().Context(), database.GetOrderForDownloadParams{
		ID:        orderID,
		CompanyID: companyID,
	})
	if err != nil {
		return nil, &errorHandling.AppError{
			Action:    "Retrieving order for PDF download",
			LogError:  fmt.Errorf("loading order %q: %w", orderID, err),
			UserError: errors.New("error: failed to retrieve the order"),
		}
	}
	if len(rows) == 0 {
		return nil, &errorHandling.AppError{
			Action:    "Retrieving order for PDF download",
			LogError:  fmt.Errorf("order %q was not found for company %d", orderID, companyID),
			UserError: errors.New("error: order not found"),
		}
	}

	first := rows[0]
	order := Order{
		ID: first.ID,
		Sender: Sender{
			Org:     first.CompanyName,
			Contact: first.PlacedBy,
		},
		Supplier: Supplier{
			Name:  first.SupplierName,
			Email: first.SupplierEmail,
		},
		Items: make([]OrderItem, 0, len(rows)),
	}
	for _, row := range rows {
		order.Items = append(order.Items, OrderItem{
			ProductName: row.NameAtOrder,
			Qty:         int(row.Quantity),
		})
	}

	pdf, aerr := os.generatePDF(order, first.CreatedAt)
	if aerr != nil {
		return nil, aerr
	}
	order.PDF = pdf

	return &order, nil
}
