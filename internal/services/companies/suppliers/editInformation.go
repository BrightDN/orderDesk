package suppliers

import (
	"database/sql"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/labstack/echo/v4"
)

func (ss *SupplierService) EditSupplier(c echo.Context, name string, compID int32, newName, email, contact, msubject, mctx string) (Supplier, error) {

	err := ss.queries.EditSupplierByNameAndCompanyID(c.Request().Context(), database.EditSupplierByNameAndCompanyIDParams{
		Name:      name,
		CompanyID: compID,
		Name_2:    newName,
		Email:     email,
		Contact:   sql.NullString{String: contact, Valid: contact != ""},
		DeletedAt: sql.NullTime{Valid: false},
	})
	if err != nil {
		return Supplier{}, err
	}
	suppl, err := ss.queries.GetSupplierByName(c.Request().Context(), newName)
	if err != nil {
		return Supplier{}, err
	}

	mail, err := ss.queries.UpdateOrderMail(c.Request().Context(), database.UpdateOrderMailParams{
		Subject:     msubject,
		MailContent: mctx,
		SupplierID:  suppl.ID,
	})
	supp := Supplier{
		ID:            suppl.ID,
		Name:          suppl.Name,
		Email:         suppl.Email,
		ContactPerson: contact,
		Active:        !suppl.DeletedAt.Valid,
		MailSubject:   mail.Subject,
		MailContext:   mail.MailContent,
	}

	return supp, nil
}
