package middlewares

import (
	"database/sql"
	"net/http"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/services/companies"
	"github.com/brightDN/orderDesk/internal/shared/session"
	"github.com/labstack/echo/v4"
)

func LoadEmployee(db *database.Queries) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			userID, ok, err := session.GetValue[int32](c, session.UserIDKey)
			if err != nil {
				return err
			}
			if !ok {
				return c.Redirect(http.StatusSeeOther, "/auth/logout")
			}
			employee, dbErr := db.GetEmployeeByUserID(c.Request().Context(), sql.NullInt32{
				Valid: true,
				Int32: userID,
			})
			if dbErr != nil {
				return c.Redirect(http.StatusSeeOther, "/auth/logout")
			}

			employeeData := companies.Employee{
				Name:       employee.DisplayName,
				Email:      employee.Email,
				Role:       employee.Role,
				EmployedAt: employee.EmployedAt,
				UserId:     int(employee.UserID.Int32),
				CompanyId:  int(employee.CompanyID),
				EmployeeId: int(employee.EmployeeID),
			}
			c.Set("employee", employeeData)
			return next(c)
		}
	}
}
