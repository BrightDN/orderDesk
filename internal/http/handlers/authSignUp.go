package handlers

import (
	"fmt"
	"html"
	"net/http"
	"net/mail"
	"time"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/brightDN/orderDesk/internal/shared/session"
	"github.com/labstack/echo/v4"
)

func (h *Handler) authSignUp(c echo.Context) error {
	token := c.Request().PostFormValue("token")
	email := c.Request().PostFormValue("email")
	_, err := mail.ParseAddress(email)
	if err != nil {
		fmt.Println("Error: Invalid email address")
		if flashErr := flash.Set(c, flash.Error, "Invalid email address"); flashErr != nil {
			return flashErr
		}
		return c.Redirect(http.StatusSeeOther, "/auth/signup/"+token)
	}
	inv, err := h.App.Services.Invitations.ValidateInvitation(c, token, email)
	if err != nil {
		fmt.Println("Error: Invalid invitation")
		if flashErr := flash.Set(c, flash.Error, err.Error()); flashErr != nil {
			return flashErr
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

	employee, err := h.App.Services.Auth.SignUp(c, email, password, name, inv)
	if err != nil {
		fmt.Println("Error: Failed to sign up user")
		if flashErr := flash.Set(c, flash.Error, err.Error()); flashErr != nil {
			return flashErr
		}
		return c.Redirect(http.StatusSeeOther, "/auth/signup/"+token)
	}
	empl, err := h.App.Db.GetEmployeeByUserID(c.Request().Context(), employee.UserID)
	if err != nil {
		fmt.Println("Error: Failed to get employee")
		if flashErr := flash.Set(c, flash.Error, err.Error()); flashErr != nil {
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
	return c.Redirect(http.StatusSeeOther, "/dashboard/neworder")
}
