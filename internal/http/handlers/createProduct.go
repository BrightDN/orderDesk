package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/brightDN/orderDesk/internal/pages"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/brightDN/orderDesk/internal/shared/logging"
	"github.com/brightDN/orderDesk/internal/shared/parse"
	"github.com/labstack/echo/v4"
)

func (h *Handler) createProduct(c echo.Context) error {
	supplier := c.Param("id")
	id, err := parse.Int32(supplier)
	if err != nil {
		if logErr := errorHandling.Log_and_flash(c, *err); logErr != nil {
			return logErr
		}
		return err.UserError
	}

	product := c.Request().PostFormValue("product")
	if strings.TrimSpace(product) == "" {
		if logErr := errorHandling.Log_and_flash(c, errorHandling.AppError{
			Action:    "Reading post form value: \"product\"",
			LogError:  fmt.Errorf("\"Product\" form data was empty, value: %s", product),
			UserError: fmt.Errorf("error: malformed request"),
		}); logErr != nil {
			return logErr
		}
		return h.returnPartialProductList(c, id)
	}

	if err = h.App.Services.Suppliers.CreateProduct(c, id, product); err != nil {
		if logErr := errorHandling.Log_and_flash(c, *err); logErr != nil {
			return logErr
		}
		return h.returnPartialProductList(c, id)
	}

	if logErr := logging.Log_info_and_flash(c, "User created a product", "Product successfully created"); logErr != nil {
		return logErr
	}
	return h.returnPartialProductList(c, id)
}

func (h *Handler) returnPartialProductList(c echo.Context, supplierID int32) error {
	pageData := pages.PageData{
		Title: "suppliers",
		Type:  pages.BusinessType,
	}
	prods, err := h.App.Services.Suppliers.GetProducts(c, supplierID)
	if err != nil {
		if logErr := errorHandling.Log_and_flash(c, *err); logErr != nil {
			return logErr
		}
		return err
	}
	supplier, err := h.App.Services.Suppliers.GetSupplierByID(c, supplierID)
	if err != nil {
		if logErr := errorHandling.Log_and_flash(c, *err); logErr != nil {
			return logErr
		}
		return err
	}
	return c.Render(http.StatusOK, "partials/itemList", map[string]any{
		"pageData": pageData,
		"products": prods,
		"supplier": supplier,
	})
}
