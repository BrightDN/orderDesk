package mailer

import (
	"fmt"

	"github.com/brightDN/orderDesk/internal/services/companies/orders"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
)

func (ms *MailerService) SendOrder(data *orders.OrderMailData) *errorHandling.AppError {
	return ms.Send(Mail{
		Receiver: data.Order.Supplier.Email,
		Subject:  data.Subject,
		Body:     data.Body,
		Attachment: &Attachment{
			Filename: fmt.Sprintf("%s.pdf", data.Order.ID),
			Reader:   data.Order.PDF,
		},
	})
}
