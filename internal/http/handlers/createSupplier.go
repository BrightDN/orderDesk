package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/brightDN/orderDesk/internal/flash"
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
		fmt.Printf("Error: Form validation failed: Name field is required, received: '%s', Email field is required, received: '%s'\n", company, email)
		if flashErr := flash.Set(c, flash.Error, ErrFormValidation.Error()); flashErr != nil {
			return flashErr
		}
		return ErrFormValidation
	}

	id, ok, err := session.GetValue[int32](c, session.CompanyIDKey)
	if err != nil {
		fmt.Printf("Error: Retrieving session information: %v", err)
		if flashErr := flash.Set(c, flash.Error, ErrSessionError.Error()); flashErr != nil {
			return flashErr
		}
		return ErrSessionError
	}
	if !ok {
		fmt.Printf("Error: Company ID not found in session: %v\n", err)
		if flashErr := flash.Set(c, flash.Error, ErrSessionError.Error()); flashErr != nil {
			return flashErr
		}
		return ErrSessionError
	}

	if err := h.App.Services.Suppliers.Create(c, company, email, contact, id); err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/app/suppliers")
}
