package handlers

import (
	"fmt"
	"html"
	"net/http"
	"strings"

	"github.com/brightDN/orderDesk/internal/http/routing"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/brightDN/orderDesk/internal/shared/session"
	"github.com/labstack/echo/v4"
)

func (h *Handler) processLogin(c echo.Context) error {
	email := c.Request().PostFormValue("email")
	password := c.Request().PostFormValue("password")

	email = html.EscapeString(email)
	password = html.EscapeString(password)

	if strings.TrimSpace(email) == "" {
		if logErr := errorHandling.Log_and_flash(c, errorHandling.AppError{
			Action:    "Processing login - email validation",
			LogError:  fmt.Errorf("Empty email field"),
			UserError: fmt.Errorf("Email is required"),
		}); logErr != nil {
			return logErr
		}
		return c.Redirect(http.StatusSeeOther, routing.Login)
	}

	if strings.TrimSpace(password) == "" {
		if logErr := errorHandling.Log_and_flash(c, errorHandling.AppError{
			Action:    "Processing login - password validation",
			LogError:  fmt.Errorf("Empty password field"),
			UserError: fmt.Errorf("Password is required"),
		}); logErr != nil {
			return logErr
		}
		return c.Redirect(http.StatusSeeOther, routing.Login)
	}

	user, appErr := h.App.Services.Auth.VerifyUser(c, email, password)
	if appErr != nil {
		if logErr := errorHandling.Log_and_flash(c, *appErr); logErr != nil {
			return logErr
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
		if logErr := errorHandling.Log_and_flash(c, errorHandling.AppError{
			Action:    "Fetching company count at login processing",
			LogError:  fmt.Errorf("Failed to get company count: %v", err),
			UserError: fmt.Errorf("Something went wrong, please try again.\nContact support if the issue persists"),
		}); logErr != nil {
			return logErr
		}
		return c.Redirect(http.StatusSeeOther, routing.Login)
	}

	if count == 1 {
		employee, err := h.App.Db.GetEmployeeByUserID(c.Request().Context(), user.ID)
		if err != nil {
			if logErr := errorHandling.Log_and_flash(c, errorHandling.AppError{
				Action:    "Fetching employee at login process",
				LogError:  fmt.Errorf("Failed to get employee: %v", err),
				UserError: fmt.Errorf("Something went wrong, please try again later or contact support"),
			}); logErr != nil {
				return logErr
			}
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
		session.SetValues(c, session.SessionData{
			UserID: user.ID,
		})
		return c.Redirect(http.StatusSeeOther, "/auth/select")
	}
}
