package companies

import (
	"github.com/labstack/echo/v4"
)

func (cs *CompanyService) GetCompanyEmployees(c echo.Context, id int32) ([]Employee, error) {
	employees, err := cs.db.GetCompanyEmployees(c.Request().Context(), id)
	if err != nil {
		return nil, ErrInternalError
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
