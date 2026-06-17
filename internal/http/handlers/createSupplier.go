package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/brightDN/orderDesk/internal/http/routing"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/brightDN/orderDesk/internal/shared/logging"
	"github.com/brightDN/orderDesk/internal/shared/session"
	"github.com/labstack/echo/v4"
)

var ErrSessionError = errors.New("error: retrieving session information failed")
var ErrFormValidation = errors.New("error: invalid form submission")

func (h *Handler) createSupplier(c echo.Context) error {
	company := c.Request().PostFormValue("company")
	email := c.Request().PostFormValue("email")
	contact := c.Request().PostFormValue("contact")

	if strings.TrimSpace(company) == "" || strings.TrimSpace(email) == "" {
		if logErr := errorHandling.Log_and_flash(c, errorHandling.AppError{
			Action:    "Reading post form values",
			LogError:  fmt.Errorf("form validation failed: Name field is required, received: '%s',\nEmail field is required, received: '%s'\n", company, email),
			UserError: ErrFormValidation,
		}); logErr != nil {
			return logErr
		}
		return redirection(c)
	}

	id, ok, err := session.GetValue[int32](c, session.CompanyIDKey)
	if err != nil {
		if handlingErr := errorHandling.Log_and_flash(c, errorHandling.AppError{
			Action:    "retrieving session data",
			LogError:  err,
			UserError: fmt.Errorf("Internal error, session terminated, please log in again"),
		}); handlingErr != nil {
			return handlingErr
		}
		return c.Redirect(http.StatusSeeOther, routing.Logout)
	}
	if !ok {
		if handlingErr := errorHandling.Log_and_flash(c, errorHandling.AppError{
			Action:    "retrieving session data",
			LogError:  fmt.Errorf("Could not retrieve companyID from the session"),
			UserError: fmt.Errorf("Internal error, session terminated, please log in again"),
		}); handlingErr != nil {
			return handlingErr
		}
		return c.Redirect(http.StatusSeeOther, routing.Logout)
	}

	if err := h.App.Services.Suppliers.Create(c, company, email, contact, id); err != nil {
		if logErr := errorHandling.Log_and_flash(c, *err); logErr != nil {
			return logErr
		}
	} else if loggingErr := logging.Log_info_and_flash(c, "User created a supplier", fmt.Sprintf("Supplier \"%s\" successfully created", company)); loggingErr != nil {
		return loggingErr
	}
	return redirection(c)
}

func redirection(c echo.Context) error {
	return c.Redirect(http.StatusSeeOther, "/app/suppliers")
}
