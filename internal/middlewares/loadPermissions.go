package middlewares

import (
	"database/sql"
	"net/http"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/services/companies"
	"github.com/brightDN/orderDesk/internal/services/permissions"
	"github.com/brightDN/orderDesk/internal/shared/session"
	"github.com/labstack/echo/v4"
)

func LoadPermissions(db *database.Queries, permissionsService *permissions.PermissionsService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			employee, employeeLoaded := c.Get("employee").(companies.Employee)
			role := employee.Role
			var customPermissions []string

			if !employeeLoaded {
				userID, userOK, err := session.GetValue[int32](c, session.UserIDKey)
				if err != nil {
					return err
				}
				companyID, companyOK, err := session.GetValue[int32](c, session.CompanyIDKey)
				if err != nil {
					return err
				}
				if !userOK || !companyOK {
					return c.Redirect(http.StatusSeeOther, "/auth/logout")
				}

				rows, dbErr := db.GetEmployeeRoleAndPermissions(c.Request().Context(), database.GetEmployeeRoleAndPermissionsParams{
					CompanyID: companyID,
					UserID:    sql.NullInt32{Int32: userID, Valid: true},
				})
				if dbErr != nil || len(rows) == 0 {
					return c.Redirect(http.StatusSeeOther, "/auth/logout")
				}

				role = rows[0].Role
				customPermissions = permissionNames(rows)
			}

			var perms permissions.Permissions
			switch role {
			case "superadmin", "admin":
			case "employee":
				perms = permissionsService.GetPermissions(role)
			case "custom":
				if employeeLoaded {
					rows, err := db.GetEmployeeRoleAndPermissions(c.Request().Context(), database.GetEmployeeRoleAndPermissionsParams{
						CompanyID: int32(employee.CompanyId),
						UserID:    sql.NullInt32{Int32: int32(employee.UserId), Valid: true},
					})
					if err != nil {
						return c.Redirect(http.StatusSeeOther, "/auth/logout")
					}
					customPermissions = permissionNames(rows)
				}
				perms = permissionsService.GetCustomPermissions(customPermissions)
			}

			c.Set(permissions.ContextKey, perms)
			return next(c)
		}
	}
}

func permissionNames(rows []database.GetEmployeeRoleAndPermissionsRow) []string {
	permissionNames := make([]string, 0, len(rows))
	for _, row := range rows {
		if row.Permission.Valid {
			permissionNames = append(permissionNames, row.Permission.String)
		}
	}
	return permissionNames
}
