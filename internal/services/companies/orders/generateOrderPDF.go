package orders

import (
	"bytes"
	"fmt"
	"time"

	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/johnfercher/maroto/v2"
)

func (o *OrderService) generatePDF(order Order, t time.Time) (*bytes.Reader, *errorHandling.AppError) {
	cfg := o.getPDFConfigs()
	m := maroto.New(cfg)

	if err := m.RegisterHeader(o.buildPDFHeader(order, t)...); err != nil {
		return nil, &errorHandling.AppError{
			Action:    "Creating order PDF",
			LogError:  fmt.Errorf("building pdf header error: %v", err),
			UserError: fmt.Errorf("error: Failed to create PDF, order has not been sent"),
		}
	}

	if err := m.RegisterFooter(o.buildPDFFooter()...); err != nil {
		return nil, &errorHandling.AppError{
			Action:    "Creating order PDF",
			LogError:  fmt.Errorf("building pdf footer error: %v", err),
			UserError: fmt.Errorf("error: Failed to create PDF, order has not been sent"),
		}
	}

	m.AddRows(o.buildPDFBody(order)...)

	document, err := m.Generate()
	if err != nil {
		return nil, &errorHandling.AppError{
			Action:    "Generating order PDF",
			LogError:  fmt.Errorf("Something went wrong generating the PDF: %v", err),
			UserError: fmt.Errorf("error: Failed to generate the email PDF, try again later or contact support if the issue persists"),
		}
	}
	return bytes.NewReader(document.GetBytes()), nil
}
