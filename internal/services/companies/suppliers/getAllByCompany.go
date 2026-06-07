package suppliers

import (
	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/labstack/echo/v4"
)

func (s *SupplierService) GetAllByCompany(c echo.Context, companyID int32) ([]Supplier, error) {
	dbSuppliers, err := s.db.GetCompanySuppliers(c.Request().Context(), companyID)
	if err != nil {
		if flashErr := flash.Set(c, flash.Error, ErrInternalError.Error()); flashErr != nil {
			return nil, flashErr
		}
		return nil, err
	}

	var suppliers []Supplier
	for _, dbSupplier := range dbSuppliers {
		contact := ""
		if dbSupplier.Contact.Valid {
			contact = dbSupplier.Contact.String
		}
		suppliers = append(suppliers, Supplier{
			ID:            dbSupplier.ID,
			Name:          dbSupplier.Name,
			Email:         dbSupplier.Email,
			ContactPerson: contact,
			Count:         dbSupplier.ProductCount,
			Active:        true,
		})
	}
	return suppliers, nil
}
