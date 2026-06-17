package suppliers

import (
	"errors"
	"fmt"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/labstack/echo/v4"
)

func (sc *SupplierService) GetSupplierByID(c echo.Context, supplierID int32) (Supplier, *errorHandling.AppError) {
	supplier, err := sc.queries.GetSupplierByID(c.Request().Context(), supplierID)
	if err != nil {
		return Supplier{}, &errorHandling.AppError{
			Action:    "Retrieving supplier by ID",
			LogError:  fmt.Errorf("Failed to fetch supplier %d: %v", supplierID, err),
			UserError: errors.New("failed to fetch supplier"),
		}
	}

	var s = Supplier{
		Name:  supplier.Name,
		Email: supplier.Email,
	}

	return s, nil
}

func (sc *SupplierService) GetSupplierByNameAndCompanyID(c echo.Context, supplierName string, companyID int32) (Supplier, *errorHandling.AppError) {
	supplier, err := sc.queries.GetCompanySupplier(c.Request().Context(), database.GetCompanySupplierParams{
		Name:      supplierName,
		CompanyID: companyID,
	})
	if err != nil {
		return Supplier{}, &errorHandling.AppError{
			Action:    "Retrieving supplier by name and company",
			LogError:  fmt.Errorf("Failed to fetch supplier %s for company %d: %v", supplierName, companyID, err),
			UserError: errors.New("failed to fetch supplier"),
		}
	}

	return Supplier{
		ID:            supplier.ID,
		Name:          supplier.Name,
		Email:         supplier.Email,
		ContactPerson: supplier.Contact.String,
		Count:         supplier.ProductCount,
		Active:        !supplier.DeletedAt.Valid,
		MailSubject:   supplier.MailSubject,
		MailContext:   supplier.MailContent,
	}, nil
}
