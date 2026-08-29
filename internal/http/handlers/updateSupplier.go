package handlers

import (
	"fmt"
	"net/http"
	"net/mail"
	"strings"

	"github.com/brightDN/orderDesk/internal/http/routing"
	"github.com/brightDN/orderDesk/internal/services/companies/suppliers"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/brightDN/orderDesk/internal/shared/logging"
	"github.com/brightDN/orderDesk/internal/shared/session"
	"github.com/labstack/echo/v4"
)

func (h *Handler) updateSupplier(c echo.Context) error {
	compID, ok, err := session.GetValue[int32](c, session.CompanyIDKey)
	if err != nil || !ok {
		if logErr := errorHandling.Log_and_flash(c, *err); logErr != nil {
			return logErr
		}
		return c.Redirect(http.StatusSeeOther, routing.Logout)
	}

	oldName := c.Param("name")
	if strings.TrimSpace(oldName) == "" {
		if logErr := errorHandling.Log_and_flash(c, errorHandling.AppError{
			Action:    "Retrieving url parameter for \"name\"",
			LogError:  fmt.Errorf("The name parameter was empty"),
			UserError: fmt.Errorf("error: malformed request, try again later or contact support"),
		}); logErr != nil {
			return logErr
		}
		return c.Redirect(http.StatusSeeOther, "/app/suppliers")
	}

	newName := c.Request().PostFormValue("name")
	if strings.TrimSpace(newName) == "" {
		appErr := errorHandling.AppError{
			Action:    "Reading formdata for updating supplier",
			LogError:  fmt.Errorf("Received an empty namefield"),
			UserError: fmt.Errorf("error: Required name field is empty"),
		}
		return renderPartialSuppInfo(c, suppliers.Supplier{}, &appErr)
	}

	email := c.Request().PostFormValue("email")
	if strings.TrimSpace(email) == "" {
		appErr := errorHandling.AppError{
			Action:    "Reading formdata for updating supplier",
			LogError:  fmt.Errorf("Received an empty emailfield"),
			UserError: fmt.Errorf("error: Required email field is empty"),
		}
		return renderPartialSuppInfo(c, suppliers.Supplier{}, &appErr)
	}

	if _, err := mail.ParseAddress(email); err != nil {
		appErr := errorHandling.AppError{
			Action:    "Validating email from post request",
			LogError:  fmt.Errorf("Received an invalid email"),
			UserError: fmt.Errorf("error: invalid email"),
		}
		return renderPartialSuppInfo(c, suppliers.Supplier{}, &appErr)
	}

	contact := c.Request().PostFormValue("contact_person")

	subject := c.Request().PostFormValue("mail_subject")
	if strings.TrimSpace(subject) == "" {
		logErr := errorHandling.AppError{
			Action:    "Validating mail subject from post form value",
			LogError:  fmt.Errorf("Retrieved an empty mail subject"),
			UserError: fmt.Errorf("error: Mail subject is empty"),
		}
		return renderPartialSuppInfo(c, suppliers.Supplier{}, &logErr)
	}
	mailCtx := c.Request().PostFormValue("order_mail")
	if strings.TrimSpace(mailCtx) == "" {
		logErr := errorHandling.AppError{
			Action:    "Validating mail subject from post form value",
			LogError:  fmt.Errorf("Retrieved an empty mail body"),
			UserError: fmt.Errorf("error: Mail body is empty"),
		}
		return renderPartialSuppInfo(c, suppliers.Supplier{}, &logErr)
	}

	supp, appErr := h.app.Services.Suppliers.EditSupplier(c, oldName, compID, newName, email, contact, subject, mailCtx)
	if appErr != nil {
		return renderPartialSuppInfo(c, suppliers.Supplier{}, appErr)
	}
	return renderPartialSuppInfo(c, supp, nil)
}

func renderPartialSuppInfo(c echo.Context, supp suppliers.Supplier, err *errorHandling.AppError) error {

	if err != nil {
		if logErr := errorHandling.Log_and_flash(c, *err); logErr != nil {
			return logErr
		}
	} else {
		if logErr := logging.Log_info_and_flash(c, "A user edited supplier info", "Supplier information succesfully updated"); logErr != nil {
			return logErr
		}
	}

	return c.Render(http.StatusOK, "partials/supplierInformation", map[string]any{
		"supplier":  supp,
		"isPartial": true,
	})
}
