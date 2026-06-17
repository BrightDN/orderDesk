package suppliers

import (
	"errors"
	"fmt"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/labstack/echo/v4"
)

func (sc *SupplierService) DeleteProduct(c echo.Context, supplierID, productID int32) *errorHandling.AppError {
	if err := sc.queries.DeleteProduct(c.Request().Context(), database.DeleteProductParams{
		SupplierID: supplierID,
		ID:         productID,
	}); err != nil {
		return &errorHandling.AppError{
			Action:    "Deleting product",
			LogError:  fmt.Errorf("Failed to delete product %d for supplier %d: %v", productID, supplierID, err),
			UserError: errors.New("failed to delete product"),
		}
	}
	return nil
}
