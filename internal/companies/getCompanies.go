package companies

import (
	"github.com/brightDN/orderDesk/internal/database"
	"github.com/labstack/echo/v4"
)

func GetCompanies(db *database.Queries, c echo.Context) []Company {
	companiesData, _ := db.GetCompanies(c.Request().Context())
	var companies = []Company{}
	for _, company := range companiesData {
		companies = append(companies, Company{
			Name:      company.Name,
			Email:     company.Email,
			CreatedAt: company.CreatedAt.Format("02-01-2006"),
			IsDeleted: company.DeletedAt.Valid,
		})
	}

	return companies
}
