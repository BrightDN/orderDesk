package companies

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/labstack/echo/v4"
)

func (cs *CompanyService) GetEmployee(c echo.Context, userID int32) (Employee, *errorHandling.AppError) {

	employee, err := cs.db.GetEmployeeByUserID(
		c.Request().Context(),
		sql.NullInt32{
			Valid: true,
			Int32: userID,
		})
	if err != nil {
		return Employee{}, &errorHandling.AppError{
			Action:    "Fetching employee details",
			LogError:  fmt.Errorf("Employee not found with user ID %d: %v", userID, err),
			UserError: errors.New("failed to fetch employee"),
		}
	}
	return Employee{
		Name:       employee.DisplayName,
		Email:      employee.Email,
		Role:       employee.Role,
		EmployedAt: employee.EmployedAt,
		UserId:     int(employee.UserID.Int32),
		CompanyId:  int(employee.CompanyID),
		EmployeeId: int(employee.EmployeeID),
	}, nil
}
