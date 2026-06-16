package handlers

import (
	"net/http"
	"net/mail"
	"strings"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/brightDN/orderDesk/internal/services/companies/suppliers"
	"github.com/brightDN/orderDesk/internal/shared/logging"
	"github.com/brightDN/orderDesk/internal/shared/session"
	"github.com/labstack/echo/v4"
)

var action = "Editing supplier"

func (h *Handler) updateSupplier(c echo.Context) error {
	compID, ok, err := session.GetValue[int32](c, session.CompanyIDKey)
	if err != nil || !ok {
		logging.ErrorLog("updating supplier", err.Error())
		if flashErr := flash.Set(c, flash.Error, err.Error()); flashErr != nil {
			return flashErr
		}
		return c.Redirect(http.StatusSeeOther, "auth/login")
	}

	oldName := c.Param("name")
	if strings.TrimSpace(oldName) == "" {
		logging.ErrorLog(action, "Received an empty name param")
		if flashErr := flash.Set(c, flash.Error, "error: malformed request"); flashErr != nil {
			return flashErr
		}
		return c.Redirect(http.StatusSeeOther, "/app/suppliers")
	}

	newName := c.Request().PostFormValue("name")
	if strings.TrimSpace(newName) == "" {
		logging.ErrorLog(action, "Required name field is empty")
		if flashErr := flash.Set(c, flash.Error, "error: name field is required"); flashErr != nil {
			return flashErr
		}
		return renderPartialSuppInfo(c, suppliers.Supplier{})
	}
	email := c.Request().PostFormValue("email")
	if strings.TrimSpace(email) == "" {
		logging.ErrorLog(action, "Required email field is empty")
		if flashErr := flash.Set(c, flash.Error, "error: email field is required"); flashErr != nil {
			return flashErr
		}
		return renderPartialSuppInfo(c, suppliers.Supplier{})
	}
	if _, err := mail.ParseAddress(email); err != nil {
		logging.ErrorLog(action, "Email is not valid")
		if flashErr := flash.Set(c, flash.Error, "error: email field is not a valid email"); flashErr != nil {
			return flashErr
		}
		return renderPartialSuppInfo(c, suppliers.Supplier{})
	}
	contact := c.Request().PostFormValue("contact_person")

	subject := c.Request().PostFormValue("mail_subject")
	if strings.TrimSpace(subject) == "" {
		logging.ErrorLog("Editing supplier", "Required subject field is empty")
		if flashErr := flash.Set(c, flash.Error, "error: subject field is required"); flashErr != nil {
			return flashErr
		}
		return renderPartialSuppInfo(c, suppliers.Supplier{})
	}
	mailCtx := c.Request().PostFormValue("order_mail")
	if strings.TrimSpace(mailCtx) == "" {
		logging.ErrorLog("Editing supplier", "Required mailcontext field is empty")
		if flashErr := flash.Set(c, flash.Error, "error: mailcontext field is required"); flashErr != nil {
			return flashErr
		}
		return renderPartialSuppInfo(c, suppliers.Supplier{})
	}

	supp, err := h.App.Services.Suppliers.EditSupplier(c, oldName, compID, newName, email, contact, subject, mailCtx)
	if err != nil {
		logging.ErrorLog("Editing supplier", err.Error())
		if flashErr := flash.Set(c, flash.Error, "Failed to update supplier"); flashErr != nil {
			logging.ErrorLog("Editing supplier", flashErr.Error())
			return c.Redirect(http.StatusSeeOther, "/app/suppliers")
		}
	}
	return renderPartialSuppInfo(c, supp)
}

func renderPartialSuppInfo(c echo.Context, supp suppliers.Supplier) error {
	return c.Render(http.StatusOK, "partials/supplierInformation", map[string]any{
		"supplier": supp,
	})
}
