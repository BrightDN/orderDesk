package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/brightDN/orderDesk/internal/http/routing"
	"github.com/brightDN/orderDesk/internal/services/authentication"
	"github.com/brightDN/orderDesk/internal/shared/session"
	"github.com/labstack/echo/v4"
)

func (h *Handler) processLogin(c echo.Context) error {
	email := c.Request().PostFormValue("email")
	password := c.Request().PostFormValue("password")

	if email == "" {
		fmt.Println("Error: email is required")
		flashErr := flash.Set(c, flash.Error, "Email is required")
		if flashErr != nil {
			return flashErr
		}
		return c.Redirect(http.StatusSeeOther, routing.Login)
	}

	if password == "" {
		fmt.Println("Error: password is required")
		flashErr := flash.Set(c, flash.Error, "Password is required")
		if flashErr != nil {
			return flashErr
		}
		return c.Redirect(http.StatusSeeOther, routing.Login)
	}

	user, err := h.App.Services.Auth.VerifyUser(c, email, password)
	if err != nil {
		fmt.Printf("Error: failed to verify user: %v\n", err)
		message := "Something went wrong, please contact support"
		if errors.Is(err, authentication.ErrInvalidCredentials) {
			message = "Invalid email or password"
		}
		flashErr := flash.Set(c, flash.Error, message)
		if flashErr != nil {
			return flashErr
		}
		return c.Redirect(http.StatusSeeOther, routing.Login)
	}

	if user.IsAdmin {
		session.SetValues(c, session.SessionData{
			UserID:         user.ID,
			IsSiteAdmin:    true,
			IsMultiCompany: false,
			RoleName:       "site_admin",
			CompanyID:      0,
		})
		return c.Redirect(http.StatusSeeOther, "/admin/companies/overview")
	}

	count, err := h.App.Db.GetCompanyCount(c.Request().Context(), user.ID)
	if err != nil || count == 0 {
		fmt.Printf("Error: failed to get company count: %v\n", err)
		flashErr := flash.Set(c, flash.Error, "Something went wrong, please contact support")
		if flashErr != nil {
			return flashErr
		}
		return c.Redirect(http.StatusSeeOther, routing.Login)
	}
	if count == 1 {
		employee, err := h.App.Db.GetEmployeeByUserID(c.Request().Context(), user.ID)
		if err != nil {
			fmt.Printf("Error: failed to get employee: %v\n", err)
			return err
		}
		session.SetValues(c, session.SessionData{
			UserID:         employee.UserID,
			CompanyID:      employee.CompanyID,
			RoleName:       employee.Role,
			IsMultiCompany: false,
		})
		return c.Redirect(http.StatusSeeOther, routing.Neworder)
	} else {
		return c.Redirect(http.StatusSeeOther, "/auth/select")
	}
}
