package suppliers

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/labstack/echo/v4"
)

func (s *SupplierService) Create(c echo.Context, company, email, contact string, companyID int32) error {
	contact = strings.TrimSpace(contact)
	_, err := s.db.CreateSupplier(c.Request().Context(), database.CreateSupplierParams{
		Name:      company,
		Email:     email,
		CompanyID: companyID,
		Contact:   sql.NullString{String: contact, Valid: contact != ""},
	})
	if err != nil {
		if flashErr := flash.Set(c, flash.Error, "Error creating supplier."); flashErr != nil {
			return flashErr
		}
		return fmt.Errorf("error creating supplier: %w", err)
	}
	return nil
}
