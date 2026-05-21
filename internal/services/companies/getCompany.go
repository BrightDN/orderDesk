package companies

import (
	"github.com/labstack/echo/v4"
)

func (cs *CompanyService) GetCompany(c echo.Context, id int32) (Company, error) {
	data, err := cs.db.GetCompany(c.Request().Context(), id)
	if err != nil {
		return Company{}, ErrInternalError
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
