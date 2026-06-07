package handlers

import (
	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/brightDN/orderDesk/internal/shared/parse"
	"github.com/labstack/echo/v4"
)

func (h *Handler) deleteProduct(c echo.Context) error {
	supplierID, err := parse.Int32(c.Param("supplierID"))
	if err != nil {
		if flashErr := flash.Trigger(c, flash.Error, err.Error()); flashErr != nil {
			return flashErr
		}
		return err
	}
	productID, err := parse.Int32(c.Param("productID"))
	if err != nil {
		if flashErr := flash.Trigger(c, flash.Error, err.Error()); flashErr != nil {
			return flashErr
		}
		return h.returnPartialProductList(c, supplierID)
	}

	if err := h.App.Services.Suppliers.DeleteProduct(c, supplierID, productID); err != nil {
		if flashErr := flash.Trigger(c, flash.Error, err.Error()); flashErr != nil {
			return flashErr
		}
		return h.returnPartialProductList(c, supplierID)
	}

	if flashErr := flash.Trigger(c, flash.Pass, "product succesfully deleted"); flashErr != nil {
		return flashErr
	}
	return h.returnPartialProductList(c, supplierID)
}
