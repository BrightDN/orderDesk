package companies

import (
	"context"
	"fmt"
	"net/http"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/labstack/echo/v4"
)

func Create(db *database.Queries, name, email string) (database.Company, error) {
	comp, err := db.CreateCompany(context.Background(), database.CreateCompanyParams{
		Name:  name,
		Email: email,
	})

	if err != nil {
		return database.Company{}, echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("Couldn't create your company: %v", err))
	}

	return comp, nil
}
