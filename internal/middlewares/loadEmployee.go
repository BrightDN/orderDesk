package middlewares

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/shared/session"
	"github.com/labstack/echo/v4"
)

func loadEmployee(db *database.Queries) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			userID, ok, err := session.GetValue[int32](c, session.UserIDKey)
			if err != nil {
				return err
			}
			if !ok {
				return c.Redirect(http.StatusSeeOther, "/auth/login")
			}
			companyID, ok, err := session.GetValue[int32](c, session.CompanyIDKey)
			if err != nil {
				return err
			}
			if !ok {
				return c.Redirect(http.StatusSeeOther, "/auth/login")
			}
			employee, err := db.GetEmployee(c.Request().Context(), database.GetEmployeeParams{
				UserID:    userID,
				CompanyID: companyID,
			})
			if err != nil {
				return c.Redirect(http.StatusSeeOther, "/auth/login")
			}
			c.Set("employee", employee)
			return next(c)
		}
	}
}
