package companies

import (
	"github.com/brightDN/orderDesk/internal/database"
	"github.com/labstack/echo/v4"
)

func (cs *CompanyService) GetEmployee(c echo.Context, companyID, userID int32) (Employee, error) {
	employee, err := cs.db.GetEmployee(c.Request().Context(), database.GetEmployeeParams{
		UserID:    userID,
		CompanyID: companyID,
	})
	if err != nil {
		return Employee{}, err
	}
	return Employee{
		Name:       employee.DisplayName,
		Email:      employee.Email,
		Role:       employee.Role,
		UserId:     int(employee.UserID),
		CompanyId:  int(employee.CompanyID),
		EmployeeId: int(employee.EmployeeID),
	}, nil
}
