package suppliers

import (
	"errors"
	"fmt"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func (sc *SupplierService) CreateProduct(c echo.Context, supplierID int32, product string) *errorHandling.AppError {
	if err := sc.queries.CreateProduct(c.Request().Context(), database.CreateProductParams{
		SupplierID: supplierID,
		Name:       product,
	}); err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return &errorHandling.AppError{
					Action:    "Creating product for supplier",
					LogError:  fmt.Errorf("Product already exists for supplier %d: %s", supplierID, product),
					UserError: errors.New("product already exists for supplier"),
				}
			}
		}
		return &errorHandling.AppError{
			Action:    "Creating product for supplier",
			LogError:  err,
			UserError: errors.New("failed to create product"),
		}
	}
	return nil
}
