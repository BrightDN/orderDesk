package handlers

import (
	"fmt"
	"html"
	"net/http"
	"net/mail"
	"time"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/brightDN/orderDesk/internal/http/routing"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/brightDN/orderDesk/internal/shared/session"
	"github.com/labstack/echo/v4"
)

func (h *Handler) authSignUp(c echo.Context) error {
	token := c.Request().PostFormValue("token")
	email := c.Request().PostFormValue("email")
	token = html.EscapeString(token)
	email = html.EscapeString(email)

	_, parseErr := mail.ParseAddress(email)
	if parseErr != nil {
		if logErr := errorHandling.Log_and_flash(c, errorHandling.AppError{
			Action:    "Signing up the user",
			LogError:  parseErr,
			UserError: fmt.Errorf("Invalid email address"),
		}); logErr != nil {
			return logErr
		}
		return c.Redirect(http.StatusSeeOther, "/auth/signup/"+token)
	}

	inv, appErr := h.App.Services.Invitations.ValidateInvitation(c, token, email)
	if appErr != nil {
		if logErr := errorHandling.Log_and_flash(c, *appErr); logErr != nil {
			return logErr
		}
		return c.Redirect(http.StatusSeeOther, "/auth/signup/"+token)
	}

	if (inv.ExpiresAt.Before(time.Now())) || inv.UsedAt.Valid {
		fmt.Println("Error: Invitation is expired or already used")
		if flashErr := flash.Set(c, flash.Error, "Invitation is not valid"); flashErr != nil {
			return flashErr
		}
		return c.Redirect(http.StatusSeeOther, "/auth/signup/"+token)
	}

	password := c.FormValue("password")
	if len(password) < 8 {
		if flashErr := flash.Set(c, flash.Error, "Password must be at least 8 characters long"); flashErr != nil {
			return flashErr
		}
		return c.Redirect(http.StatusSeeOther, "/auth/signup/"+token)
	}

	name := html.EscapeString(c.FormValue("name"))
	if len(name) < 2 {
		if flashErr := flash.Set(c, flash.Error, "Name must be at least 2 characters long"); flashErr != nil {
			return flashErr
		}
		return c.Redirect(http.StatusSeeOther, "/auth/signup/"+token)
	}

	employee, appErr := h.App.Services.Auth.SignUp(c, email, password, name, inv)
	if appErr != nil {
		fmt.Println("Error: Failed to sign up user")
		if flashErr := flash.Set(c, flash.Error, appErr.UserError.Error()); flashErr != nil {
			return flashErr
		}
		return c.Redirect(http.StatusSeeOther, "/auth/signup/"+token)
	}
	empl, dbErr := h.App.Db.GetEmployeeByUserID(c.Request().Context(), employee.UserID)
	if dbErr != nil {
		fmt.Println("Error: Failed to get employee")
		if flashErr := flash.Set(c, flash.Error, dbErr.Error()); flashErr != nil {
			return flashErr
		}
		return c.Redirect(http.StatusSeeOther, "/auth/signup/"+token)
	}

	if err := session.SetValues(c, session.SessionData{
		UserID:         empl.UserID,
		RoleName:       empl.Role,
		CompanyID:      employee.CompanyID,
		IsMultiCompany: false}); err != nil {
		fmt.Println("Error: Failed to set session values")
		if flashErr := flash.Set(c, flash.Error, err.Error()); flashErr != nil {
			return flashErr
		}
		return c.Redirect(http.StatusSeeOther, "/auth/signup/"+token)
	}
	return c.Redirect(http.StatusSeeOther, routing.NeworderPage)
}
