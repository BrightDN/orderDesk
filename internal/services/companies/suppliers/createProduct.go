package suppliers

import (
	"errors"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

var ErrSupplierProductNotUnique = errors.New("Error: product already exists for supplier")

func (sc *SupplierService) CreateProduct(c echo.Context, supplierID int32, product string) error {
	if err := sc.db.CreateProduct(c.Request().Context(), database.CreateProductParams{
		SupplierID: supplierID,
		Name:       product,
	}); err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return ErrSupplierProductNotUnique
			}
		}
		return ErrInternalError
	}
	return nil
}
