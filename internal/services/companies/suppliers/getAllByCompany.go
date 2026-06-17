package suppliers

import (
	"errors"
	"fmt"

	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/labstack/echo/v4"
)

func (s *SupplierService) GetAllByCompany(c echo.Context, companyID int32) ([]Supplier, *errorHandling.AppError) {
	dbSuppliers, err := s.queries.GetCompanySuppliers(c.Request().Context(), companyID)
	if err != nil {
		return nil, &errorHandling.AppError{
			Action:    "Fetching company suppliers",
			LogError:  fmt.Errorf("Failed to fetch suppliers for company %d: %v", companyID, err),
			UserError: errors.New("failed to fetch suppliers"),
		}
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
			Active:        !dbSupplier.DeletedAt.Valid,
			MailSubject:   dbSupplier.MailSubject,
			MailContext:   dbSupplier.MailContent,
		})
	}
	return suppliers, nil
}
