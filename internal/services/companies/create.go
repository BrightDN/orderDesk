package companies

import (
	"errors"
	"fmt"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

var ErrDuplicateName = errors.New("this name already exists")
var ErrDuplicateEmail = errors.New("this email already exists")

func Create(db *database.Queries, c echo.Context, name, email string) (database.Company, error) {
	comp, err := db.CreateCompany(c.Request().Context(), database.CreateCompanyParams{
		Name:  name,
		Email: email,
	})

	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) &&
			pqErr.Code == "23505" {
			switch pqErr.Constraint {
			case "unique_name":
				return database.Company{}, fmt.Errorf("%w", ErrDuplicateName)
			case "unique_email":
				return database.Company{}, fmt.Errorf("%w", ErrDuplicateEmail)
			}
		}
	}
	return comp, nil
}
