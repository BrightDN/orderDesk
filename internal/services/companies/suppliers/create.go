package suppliers

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/lib/pq"

	"github.com/labstack/echo/v4"
)

func (s *SupplierService) Create(c echo.Context, company, email, contact string, companyID int32) *errorHandling.AppError {
	contact = strings.TrimSpace(contact)
	tx, err := s.db.BeginTx(c.Request().Context(), nil)
	if err != nil {
		return &errorHandling.AppError{
			Action:    "Creating supplier - beginning transaction",
			LogError:  err,
			UserError: errors.New("failed to create supplier"),
		}
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
					return &errorHandling.AppError{
						Action:    "Creating supplier",
						LogError:  fmt.Errorf("Duplicate supplier name: %s", company),
						UserError: errors.New("a supplier with that name already exists"),
					}
				case "suppliers_company_email_unique":
					return &errorHandling.AppError{
						Action:    "Creating supplier",
						LogError:  fmt.Errorf("Duplicate supplier email: %s", email),
						UserError: errors.New("a supplier with that email already exists"),
					}
				}
			}
		}
		return &errorHandling.AppError{
			Action:    "Creating supplier",
			LogError:  err,
			UserError: errors.New("failed to create supplier"),
		}
	}

	if err := queries.CreateOrderMail(c.Request().Context(), supp.ID); err != nil {
		return &errorHandling.AppError{
			Action:    "Creating supplier - adding order mail template",
			LogError:  err,
			UserError: errors.New("failed to create supplier"),
		}
	}

	if err := tx.Commit(); err != nil {
		return &errorHandling.AppError{
			Action:    "Creating supplier - committing transaction",
			LogError:  err,
			UserError: errors.New("failed to create supplier"),
		}
	}

	return nil
}
