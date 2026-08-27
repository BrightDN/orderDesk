package handlers

import (
	"fmt"
	"net/mail"

	"github.com/brightDN/orderDesk/internal/services/companies"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/labstack/echo/v4"
)

func (h *Handler) inviteEmployee(c echo.Context) error {
	employee := c.Get("employee").(companies.Employee)
	email := c.Request().PostFormValue("email")
	if !valid(email) {
		errorHandling.Log_and_flash(c, errorHandling.AppError{
			Action:    "Validating email from post request -- adding employee",
			LogError:  fmt.Errorf("No valid emailaddress was provided"),
			UserError: fmt.Errorf("The given emailaddress is invalid"),
		})
	}

	if employee.Role != "superadmin" || employee.Role != "admin" {
		errorHandling.Log_and_flash(c, errorHandling.AppError{
			Action:    "Validating logged in users permissions",
			LogError:  fmt.Errorf("No valid emailaddress was provided"),
			UserError: fmt.Errorf("The given emailaddress is invalid"),
		})
	}
}

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
