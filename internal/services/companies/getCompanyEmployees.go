package companies

import (
	"errors"
	"fmt"

	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/labstack/echo/v4"
)

func (cs *CompanyService) GetCompanyEmployees(c echo.Context, id int32) ([]Employee, *errorHandling.AppError) {
	employees, err := cs.db.GetCompanyEmployees(c.Request().Context(), id)
	if err != nil {
		return nil, &errorHandling.AppError{
			Action:    "Fetching company employees",
			LogError:  fmt.Errorf("Failed to fetch employees for company %d: %v", id, err),
			UserError: errors.New("failed to fetch employees"),
		}
	}

	var cEmployees = []Employee{}
	for _, employee := range employees {
		cEmployees = append(cEmployees, Employee{
			Name:       employee.DisplayName,
			Email:      employee.Email,
			Role:       employee.Role,
			UserId:     int(employee.UserID),
			CompanyId:  int(employee.CompanyID),
			EmployeeId: int(employee.EmployeeID),
		})
	}

	return cEmployees, nil
}
