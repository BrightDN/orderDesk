package orders

import (
	"time"

	"github.com/brightDN/orderDesk/internal/services/companies"
	"github.com/brightDN/orderDesk/internal/services/companies/suppliers"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
)

func (os *OrderService) GetOrderData(response OrderResponse, supplier suppliers.Supplier, employee companies.Employee) (*Order, *errorHandling.AppError) {
	now := time.Now()
	orderData := os.setOrderData(supplier, employee, response, now)
	PDFBytes, err := os.generatePDF(orderData, now)
	if err != nil {
		return nil, err
	}

	orderData.PDF = PDFBytes

	return &orderData, nil
}
