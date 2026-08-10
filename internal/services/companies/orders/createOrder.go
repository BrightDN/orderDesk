package orders

import (
	"encoding/json"
	"errors"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/labstack/echo/v4"
)

type bulkOrderItem struct {
	OrderID     string `json:"order_id"`
	ProductID   int32  `json:"product_id"`
	Quantity    int    `json:"quantity"`
	NameAtOrder string `json:"name_at_order"`
}

func (os *OrderService) CreateOrder(c echo.Context, companyID, supplierID, employeeID int32, orderData *Order) *errorHandling.AppError {
	if len(orderData.Items) == 0 {
		return &errorHandling.AppError{
			Action:    "Creating order",
			LogError:  errors.New("order has no items"),
			UserError: errors.New("error: an order must contain at least one item"),
		}
	}

	tx, err := os.db.BeginTx(c.Request().Context(), nil)
	if err != nil {
		return &errorHandling.AppError{
			Action:    "Beginning database transaction for creating order",
			LogError:  err,
			UserError: errors.New("error: failed to create order"),
		}
	}
	defer tx.Rollback()
	queries := database.New(tx)

	if err = queries.CreateOrder(c.Request().Context(), database.CreateOrderParams{
		ID:         orderData.ID,
		SupplierID: supplierID,
		CompanyID:  companyID,
		EmployeeID: employeeID,
	}); err != nil {
		return &errorHandling.AppError{
			Action:    "Creating order",
			LogError:  err,
			UserError: errors.New("error: failed to create order"),
		}
	}

	products, err := queries.GetProducts(c.Request().Context(), supplierID)
	if err != nil {
		return &errorHandling.AppError{
			Action:    "Loading supplier products for order",
			LogError:  err,
			UserError: errors.New("error: failed to create order"),
		}
	}

	productIDs := make(map[string]int32, len(products))
	for _, product := range products {
		productIDs[product.Name] = product.ID
	}

	items := make([]bulkOrderItem, 0, len(orderData.Items))
	for _, item := range orderData.Items {
		productID, ok := productIDs[item.ProductName]
		if !ok {
			return &errorHandling.AppError{
				Action:    "Matching order item to supplier product",
				LogError:  errors.New("supplier product not found: " + item.ProductName),
				UserError: errors.New("error: one of the selected products is no longer available"),
			}
		}

		items = append(items, bulkOrderItem{
			OrderID:     orderData.ID,
			ProductID:   productID,
			Quantity:    item.Qty,
			NameAtOrder: item.ProductName,
		})
	}

	payload, err := json.Marshal(items)
	if err != nil {
		return &errorHandling.AppError{
			Action:    "Encoding order items for bulk insert",
			LogError:  err,
			UserError: errors.New("error: failed to create order"),
		}
	}

	if _, err = queries.CreateOrderItems(c.Request().Context(), json.RawMessage(payload)); err != nil {
		return &errorHandling.AppError{
			Action:    "Bulk creating order items",
			LogError:  err,
			UserError: errors.New("error: failed to create order items"),
		}
	}

	if err = tx.Commit(); err != nil {
		return &errorHandling.AppError{
			Action:    "Committing transaction for creating order",
			LogError:  err,
			UserError: errors.New("error: failed to create order"),
		}
	}
	return nil
}
