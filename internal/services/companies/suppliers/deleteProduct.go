package suppliers

import (
	"fmt"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/labstack/echo/v4"
)

func (sc *SupplierService) DeleteProduct(c echo.Context, supplierID, productID int32) error {
	if err := sc.queries.DeleteProduct(c.Request().Context(), database.DeleteProductParams{
		SupplierID: supplierID,
		ID:         productID,
	}); err != nil {
		fmt.Printf("Error: Failed to delete product: %v", err)
		return ErrInternalError
	}
	return nil
}
