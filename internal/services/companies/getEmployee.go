package companies

import (
	"github.com/labstack/echo/v4"
)

func (cs *CompanyService) GetEmployee(c echo.Context, userID int32) (Employee, error) {
	employee, err := cs.db.GetEmployeeByUserID(c.Request().Context(), userID)
	if err != nil {
		return Employee{}, err
	}
	return Employee{
		Name:       employee.DisplayName,
		Email:      employee.Email,
		Role:       employee.Role,
		EmployedAt: employee.EmployedAt,
		UserId:     int(employee.UserID),
		CompanyId:  int(employee.CompanyID),
		EmployeeId: int(employee.EmployeeID),
	}, nil
}
