package suppliers

import (
	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/labstack/echo/v4"
)

func (sc *SupplierService) GetProducts(c echo.Context, id int32) ([]Products, error) {
	prodsDB, err := sc.queries.GetProducts(c.Request().Context(), id)
	if err != nil {
		if flashErr := flash.Set(c, flash.Error, ErrInternalError.Error()); flashErr != nil {
			return nil, flashErr
		}
		return nil, err
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

type Products struct {
	ID         int32
	SupplierID int32
	Name       string
}
