package orders

import (
	"time"

	"github.com/brightDN/orderDesk/internal/services/companies"
	"github.com/brightDN/orderDesk/internal/services/companies/suppliers"
)

func (os *OrderService) setOrderData(su suppliers.Supplier, se companies.Employee, or OrderResponse, t time.Time) Order {

	id := os.generateOrderID(t, su.ID)
	return Order{
		ID: id,
		Sender: Sender{
			Org:     se.EmployedAt,
			Contact: se.Name,
		},
		Supplier: Supplier{
			Name:  su.Name,
			Email: su.Email,
		},
		Items: or.Items,
	}
}
