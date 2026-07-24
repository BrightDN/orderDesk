package orders

import (
	"bytes"

	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

var (
	// colorBrand is the green used by the OrderDesk logo mark (#16A34A).
	colorBrand = &props.Color{Red: 22, Green: 163, Blue: 74}

	// colorAccent is the teal-green used for the selected-supplier card border (#1D9E75).
	colorAccent = &props.Color{Red: 29, Green: 158, Blue: 117}

	// colorInk is the near-black used for headings and strong text (#1A1A18).
	colorInk = &props.Color{Red: 26, Green: 26, Blue: 24}

	// colorMuted is the soft gray used for uppercase labels and secondary text (#8E908A).
	colorMuted = &props.Color{Red: 142, Green: 144, Blue: 138}

	// colorLink is the blue used for email addresses (#1A6DBC).
	colorLink = &props.Color{Red: 26, Green: 109, Blue: 188}

	// colorBorder is the light gray used for card borders and table dividers (#E4E4E2).
	colorBorder = &props.Color{Red: 228, Green: 228, Blue: 226}
)

func (o *OrderService) getPDFConfigs() *entity.Config {
	return config.NewBuilder().
		WithOrientation(orientation.Vertical).
		WithPageNumber(props.PageNumber{
			Pattern: "Page {current} of {total}",
			Place:   props.Bottom,
			Size:    7,
			Color:   colorMuted,
		}).
		WithPageSize(pagesize.A4).
		WithLeftMargin(15).
		WithRightMargin(15).
		WithTopMargin(12).
		WithBottomMargin(12).
		Build()
}

type OrderResponse struct {
	SupplierName string      `json:"supplierName"`
	Items        []OrderItem `json:"items"`
}

type OrderItem struct {
	ProductName string `json:"productName"`
	Qty         int    `json:"qty"`
}

// TotalQuantity adds up every line on the order.
func (o Order) totalQuantity() int {
	total := 0
	for _, item := range o.Items {
		total += item.Qty
	}
	return total
}

type Order struct {
	PDF      *bytes.Reader
	ID       string
	Sender   Sender
	Supplier Supplier
	Items    []OrderItem
}

// Sender is the company/user placing the order (the "from" side).
type Sender struct {
	Org     string
	Contact string
}

// Supplier is the company the order is sent to (the "to" side).
type Supplier struct {
	Name  string
	Email string
}
