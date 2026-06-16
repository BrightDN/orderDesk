package suppliers

import (
	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/shared/logging"
	"github.com/labstack/echo/v4"
)

const ACTION_GETSUPPLIER = "Retrieving supplier"

func (sc *SupplierService) GetSupplierByID(c echo.Context, supplierID int32) (Supplier, error) {
	supplier, err := sc.queries.GetSupplierByID(c.Request().Context(), supplierID)
	if err != nil {
		return Supplier{}, ErrInternalError
	}

	var s = Supplier{
		Name:  supplier.Name,
		Email: supplier.Email,
	}

	return s, nil
}

func (sc *SupplierService) GetSupplierByNameAndCompanyID(c echo.Context, supplierName string, companyID int32) (Supplier, error) {
	supplier, err := sc.queries.GetCompanySupplier(c.Request().Context(), database.GetCompanySupplierParams{
		Name:      supplierName,
		CompanyID: companyID,
	})
	if err != nil {
		logging.ErrorLog(ACTION_GETSUPPLIER, err.Error())
		return Supplier{}, ErrInternalError
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
