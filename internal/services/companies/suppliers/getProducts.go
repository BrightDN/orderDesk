package suppliers

import (
	"errors"
	"fmt"

	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/labstack/echo/v4"
)

type Products struct {
	ID         int32
	SupplierID int32
	Name       string
}

func (sc *SupplierService) GetProducts(c echo.Context, id int32) ([]Products, *errorHandling.AppError) {
	prodsDB, err := sc.queries.GetProducts(c.Request().Context(), id)
	if err != nil {
		return nil, &errorHandling.AppError{
			Action:    "Fetching supplier products",
			LogError:  fmt.Errorf("Failed to fetch products for supplier %d: %v", id, err),
			UserError: errors.New("failed to fetch products"),
		}
	}

	var prods = []Products{}
	for _, prod := range prodsDB {
		prods = append(prods, Products{
			ID:         prod.ID,
			SupplierID: prod.SupplierID,
			Name:       prod.Name,
		})
	}
	return prods, nil
}
