package suppliers

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/brightDN/orderDesk/internal/shared/logging"
	"github.com/lib/pq"

	"github.com/labstack/echo/v4"
)

var action = "create supplier"

var ErrSupplierNameExists = errors.New("a supplier with that name already exists")
var ErrSupplierEmailExists = errors.New("a supplier with that email already exists")

func (s *SupplierService) Create(c echo.Context, company, email, contact string, companyID int32) error {
	contact = strings.TrimSpace(contact)
	tx, err := s.db.BeginTx(c.Request().Context(), nil)
	if err != nil {
		logging.ErrorLog(action, "Failed to begin database transaction")
		if flashErr := flash.Set(c, flash.Error, ErrInternalError.Error()); flashErr != nil {
			return flashErr
		}
		return ErrInternalError
	}
	defer tx.Rollback()

	queries := database.New(tx)

	supp, err := queries.CreateSupplier(c.Request().Context(), database.CreateSupplierParams{
		Name:      company,
		Email:     email,
		CompanyID: companyID,
		Contact:   sql.NullString{String: contact, Valid: contact != ""},
	})
	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				switch pgErr.Constraint {
				case "suppliers_company_name_unique":
					return ErrSupplierNameExists
				case "suppliers_company_email_unique":
					return ErrSupplierEmailExists
				}
			}
		}
		if flashErr := flash.Set(c, flash.Error, ErrInternalError.Error()); flashErr != nil {
			return flashErr
		}
		return ErrInternalError
	}

	if err := queries.CreateOrderMail(c.Request().Context(), supp.ID); err != nil {
		logging.ErrorLog(action, "Failed to insert order mail into the db")
		if flashErr := flash.Set(c, flash.Error, ErrInternalError.Error()); flashErr != nil {
			return flashErr
		}
		return ErrInternalError
	}

	if err := tx.Commit(); err != nil {
		logging.ErrorLog(action, "Failed to commit transaction")
		if flashErr := flash.Set(c, flash.Error, ErrInternalError.Error()); flashErr != nil {
			return flashErr
		}
		return ErrInternalError
	}

	return nil
}
