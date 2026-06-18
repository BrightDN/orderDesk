package handlers

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/brightDN/orderDesk/internal/shared/logging"
	"github.com/brightDN/orderDesk/internal/shared/parse"
	"github.com/labstack/echo/v4"
)

func (h *Handler) deleteProduct(c echo.Context) error {
	supplierID, err := parse.Int32(c.Param("supplierID"))
	if err != nil {
		if logErr := errorHandling.Log_and_flash(c, *err); logErr != nil {
			return logErr
		}
		return c.NoContent(http.StatusBadRequest)
	}
	productID, err := parse.Int32(c.Param("productID"))
	if err != nil {
		if logErr := errorHandling.Log_and_flash(c, *err); logErr != nil {
			return logErr
		}
		return h.returnPartialProductList(c, supplierID)
	}

	if err := h.App.Services.Suppliers.DeleteProduct(c, supplierID, productID); err != nil {
		if logErr := errorHandling.Log_and_flash(c, *err); logErr != nil {
			return logErr
		}
		return h.returnPartialProductList(c, supplierID)
	}

	if logErr := logging.Log_info_and_flash(c, "User deleted a product", "Product successfully deleted"); logErr != nil {
		return logErr
	}
	return h.returnPartialProductList(c, supplierID)
}
