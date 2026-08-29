package handlers

import (
	"fmt"
	"net/mail"

	"github.com/brightDN/orderDesk/internal/services/permissions"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/labstack/echo/v4"
)

func (h *Handler) inviteEmployee(c echo.Context) error {
	email := c.Request().PostFormValue("email")
	if !valid(email) {
		errorHandling.Log_and_flash(c, errorHandling.AppError{
			Action:    "Validating email from post request -- adding employee",
			LogError:  fmt.Errorf("No valid emailaddress was provided"),
			UserError: fmt.Errorf("The given emailaddress is invalid"),
		})
	}
	permissions := c.Get("permissions").(permissions.Permissions)
	if !permissions.CanEditCompany {
		errorHandling.Log_and_flash(c, errorHandling.AppError{
			Action:    "Validating logged in users permissions",
			LogError:  fmt.Errorf("this user is not autorized to edit the company."),
			UserError: fmt.Errorf("You do not have the rights to perform this action."),
		})
	}

	return nil
}

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
