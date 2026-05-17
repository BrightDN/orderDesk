package companies

import (
	"github.com/brightDN/orderDesk/internal/database"
	"github.com/labstack/echo/v4"
)

func GetCompany(db *database.Queries, c echo.Context) (Company, error) {
	id, err := convertIDParam(c)
	if err != nil {
		return Company{}, err
	}

	data, err := db.GetCompany(c.Request().Context(), int32(id))
	if err != nil {
		return Company{}, err
	}

	company := Company{
		ID:        data.ID,
		Name:      data.Name,
		Email:     data.Email,
		CreatedAt: data.CreatedAt.Format("02-01-2006"),
		IsDeleted: data.DeletedAt.Valid,
	}

	return company, nil
}
