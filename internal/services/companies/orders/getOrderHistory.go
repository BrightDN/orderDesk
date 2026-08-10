package orders

import (
	"fmt"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/labstack/echo/v4"
)

func (os *OrderService) GetOrderHistoryForSupplier(c echo.Context, supplierID, companyID int32) ([]OrderHistoryLog, *errorHandling.AppError) {
	h, err := os.queries.GetOrdersBySupplier(c.Request().Context(), database.GetOrdersBySupplierParams{
		SupplierID: supplierID,
		CompanyID:  companyID,
	})
	if err != nil {
		aerr := errorHandling.AppError{
			Action:    "Retrieving order history for supplier",
			LogError:  fmt.Errorf("something went wrong during the database call: %v", err),
			UserError: fmt.Errorf("error: failed to retrieve information from the server. Try again later or contact support."),
		}
		return []OrderHistoryLog{}, &aerr
	}

	var hLogs []OrderHistoryLog
	for _, hLog := range h {
		hLogs = append(hLogs, OrderHistoryLog{
			OrderID:   hLog.ID,
			Date:      hLog.CreatedAt.Format("02 Jan 2006"),
			Time:      hLog.CreatedAt.Format("15:04"),
			Placed_by: hLog.PlacedBy,
			ItemCount: int(hLog.ItemCount),
		})
	}

	return hLogs, nil
}

type OrderHistoryLog struct {
	ItemCount int
	OrderID   string
	Placed_by string
	Date      string
	Time      string
}
