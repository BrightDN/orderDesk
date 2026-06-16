package handlers

import (
	"net/http"
	"strings"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/brightDN/orderDesk/internal/pages"
	"github.com/brightDN/orderDesk/internal/shared/parse"
	"github.com/labstack/echo/v4"
)

func (h *Handler) createProduct(c echo.Context) error {
	supplier := c.Param("id")
	id, err := parse.Int32(supplier)
	if err != nil {
		if flashErr := flash.Trigger(c, flash.Error, err.Error()); flashErr != nil {
			return flashErr
		}
		return err
	}

	product := c.Request().PostFormValue("product")
	if strings.TrimSpace(product) == "" {
		if flashErr := flash.Trigger(c, flash.Error, ErrFormValidation.Error()); flashErr != nil {
			return flashErr
		}
		return h.returnPartialProductList(c, id)
	}

	if err = h.App.Services.Suppliers.CreateProduct(c, id, product); err != nil {
		if flashErr := flash.Trigger(c, flash.Error, err.Error()); flashErr != nil {
			return flashErr
		}
		return h.returnPartialProductList(c, id)
	}

	if flashErr := flash.Trigger(c, flash.Pass, "product created"); flashErr != nil {
		return flashErr
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
		return err
	}
	supplier, err := h.App.Services.Suppliers.GetSupplierByID(c, supplierID)
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, "partials/itemList", map[string]any{
		"pageData": pageData,
		"products": prods,
		"supplier": supplier,
	})
}
