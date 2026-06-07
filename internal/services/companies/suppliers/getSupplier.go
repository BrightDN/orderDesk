package suppliers

import "github.com/labstack/echo/v4"

func (sc *SupplierService) GetSupplier(c echo.Context, supplierID int32) (Supplier, error) {
	supplier, err := sc.db.GetSupplierByID(c.Request().Context(), supplierID)
	if err != nil {
		return Supplier{}, ErrInternalError
	}

	var s = Supplier{
		Name:  supplier.Name,
		Email: supplier.Email,
	}

	return s, nil
}
