package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/labstack/echo/v4"
)

func (h *Handler) sendOrder(c echo.Context) error {

	var order OrderResponse

	if err := c.Bind(&order); err != nil {
		return err
	}

	// Generate PDF
	if err := generateOrderPDF(sampleOrder()); err != nil {
		log.Printf("PDF generation error: %v\n", err)
	}

	return returnFeedback(c, flash.Pass, "Order has been sent")
}

func returnFeedback(c echo.Context, t flash.MessageType, msg string) error {
	feedback := flash.Flash{
		Type:    t,
		Message: msg,
	}
	return c.Render(http.StatusOK, "/components/feedback", map[string]any{
		"feedback": feedback,
	})
}

func generateOrderPDF(order Order) error {
	cfg := config.NewBuilder().
		WithOrientation(orientation.Vertical).
		WithPageSize(pagesize.A4).
		WithLeftMargin(15).
		WithRightMargin(15).
		WithTopMargin(12).
		WithBottomMargin(12).
		Build()

	m := maroto.New(cfg)

	if err := m.RegisterHeader(buildHeader(order)...); err != nil {
		log.Fatal(err)
	}
	if err := m.RegisterFooter(buildFooter()...); err != nil {
		log.Fatal(err)
	}

	m.AddRows(buildBody(order)...)

	// Generate and save PDF
	document, err := m.Generate()
	if err != nil {
		return fmt.Errorf("generate PDF: %w", err)
	}

	timestamp := time.Now().Format("20060102_150405")
	filePath := fmt.Sprintf("assets/pdfs/order_%s.pdf", timestamp)
	if err := document.Save(filePath); err != nil {
		return fmt.Errorf("save PDF to %s: %w", filePath, err)
	}

	log.Printf("Order PDF saved to %s\n", filePath)
	return nil
}

// CORRECTED DATA

type OrderResponse struct {
	SupplierID int         `json:"supplierID"`
	Items      []OrderItem `json:"items"`
}

type OrderItem struct {
	ProductName string `json:"productName"`
	Qty         int    `json:"qty"`
}

// TotalQuantity adds up every line on the order.
func (o Order) TotalQuantity() int {
	total := 0
	for _, item := range o.Items {
		total += item.Qty
	}
	return total
}

// TEMPORAL TEST DATA

type Order struct {
	OrderID  string
	Date     time.Time
	Sender   Sender
	Supplier Supplier
	Items    []OrderItem
}

// Sender is the company/user placing the order (the "from" side).
type Sender struct {
	Org     string
	Contact string
	Role    string
}

// Supplier is the company the order is sent to (the "to" side).
type Supplier struct {
	Name         string
	Email        string
	Status       string
	ProductCount int
}

// sampleOrder mirrors the "New Order — Inex23" screen: comp admin at
// OrderDesk sending second/third/fourth to Inex23 <in@ex.co>.
// Swap this out for real data (e.g. loaded from your DB/API) in production.
func sampleOrder() Order {
	now := time.Now()
	orderID := generateOrderID(now)
	return Order{
		OrderID: orderID,
		Date:    now,
		Sender: Sender{
			Org:     "Proxy Wieze",
			Contact: "B. De Neef",
			Role:    "superadmin",
		},
		Supplier: Supplier{
			Name:  "Inex23",
			Email: "in@ex.co",
		},
		Items: []OrderItem{
			{ProductName: "1L Volle melk brik", Qty: 2},
			{ProductName: "1L Volle melk plastiek", Qty: 9},
			{ProductName: "0.5L Halfvolle melk plastiek", Qty: 20},
			{ProductName: "Karnemelk", Qty: 1},
			{ProductName: "Yoghurt natuur", Qty: 1},
			{ProductName: "Yoghurt natuur vol", Qty: 5},
			{ProductName: "Yoghurt natuur vol", Qty: 5},
			{ProductName: "Yoghurt natuur vol", Qty: 5},
			{ProductName: "Yoghurt natuur vol", Qty: 5},
			{ProductName: "Yoghurt natuur vol", Qty: 5},
			{ProductName: "Yoghurt natuur vol", Qty: 5},
			{ProductName: "Yoghurt natuur vol", Qty: 5},
			{ProductName: "Yoghurt natuur vol", Qty: 5},
			{ProductName: "Yoghurt natuur vol", Qty: 5},
			{ProductName: "Yoghurt natuur vol", Qty: 5},
			{ProductName: "Yoghurt natuur vol", Qty: 5},
			{ProductName: "Yoghurt natuur vol", Qty: 5},
			{ProductName: "Yoghurt natuur vol", Qty: 5},
			{ProductName: "Yoghurt natuur vol", Qty: 5},
			{ProductName: "Yoghurt natuur vol", Qty: 5},
			{ProductName: "Yoghurt natuur vol", Qty: 5},
			{ProductName: "Yoghurt natuur vol", Qty: 5},
			{ProductName: "Yoghurt natuur vol", Qty: 5},
			{ProductName: "Yoghurt natuur vol", Qty: 5},
		},
	}
}

func generateOrderID(currentTime time.Time) string {
	return "ORD-" + currentTime.Format("20060102-1504")
}

// buildHeader renders the OrderDesk wordmark + logo on the left and the
// "PURCHASE ORDER" meta block on the right, capped with the accent rule.
// It is registered with RegisterHeader, so it repeats on every page.
func buildHeader(o Order) []core.Row {
	return []core.Row{
		row.New(20).Add(
			image.NewFromFileCol(6, "assets/images/logo/orderdesk_logo.png", props.Rect{
				Center: true,
			}),

			col.New(5).Add(
				text.New("PURCHASE ORDER", props.Text{Size: 7, Style: fontstyle.Bold, Align: align.Right, Color: colorMuted, Top: 1}),
				text.New(o.OrderID, props.Text{Size: 11, Style: fontstyle.Bold, Align: align.Right, Color: colorInk, Top: 5.5}),
				text.New(o.Date.Format("Jan 2, 2006"), props.Text{Size: 7, Align: align.Right, Color: colorMuted, Top: 11}),
			),
		),
		row.New(3),
		row.New(1).Add(
			line.NewCol(12, props.Line{Thickness: 0.7, Color: colorAccent, SizePercent: 100}),
		),
		row.New(4),
	}
}

// buildFooter renders a thin rule and a small "generated by" credit line.
// It is registered with RegisterFooter, so it repeats on every page.
func buildFooter() []core.Row {
	return []core.Row{
		row.New(1).Add(
			line.NewCol(12, props.Line{Thickness: 0.3, Color: colorBorder, SizePercent: 100}),
		),
		row.New(8).Add(
			col.New(8).Add(
				text.New("Generated by OrderDesk", props.Text{Size: 7, Color: colorMuted, Top: 3}),
			),
			col.New(4).Add(
				text.New("© all rights reserved to OrderDesk", props.Text{Size: 7, Align: align.Right, Color: colorMuted, Top: 3}),
			),
		),
	}
}

// buildBody renders the page title, the from/to cards, the products table
// and the "sent" confirmation stamp. It is added with AddRows.
func buildBody(o Order) []core.Row {
	rows := []core.Row{
		row.New(4),

		// Title, matching the "New Order" heading on screen.
		row.New(12).Add(
			col.New(12).Add(
				text.New("New Order", props.Text{Size: 18, Style: fontstyle.Bold, Color: colorInk}),
			),
		),
		row.New(3),

		// "Order from" / "Order to" cards. The right-hand card reuses the
		// accent green border, the same way the selected supplier card is
		// outlined in green on screen.
		row.New(26).Add(
			col.New(5).WithStyle(&props.Cell{
				BorderColor:     colorBorder,
				BorderType:      border.Full,
				BorderThickness: 0.3,
			}).Add(
				text.New("ORDER FROM", props.Text{Size: 7, Style: fontstyle.Bold, Color: colorMuted, Top: 3, Left: 4}),
				text.New(o.Sender.Org, props.Text{Size: 10, Style: fontstyle.Bold, Color: colorInk, Top: 9, Left: 4}),
				text.New(fmt.Sprintf("Placed by %s", o.Sender.Contact), props.Text{Size: 8, Color: colorMuted, Top: 15, Left: 4}),
			),
			col.New(2),
			col.New(5).WithStyle(&props.Cell{
				BorderColor:     colorAccent,
				BorderType:      border.Full,
				BorderThickness: 0.4,
			}).Add(
				text.New("ORDER TO", props.Text{Size: 7, Style: fontstyle.Bold, Color: colorMuted, Top: 3, Left: 4}),
				text.New(o.Supplier.Name, props.Text{Size: 10, Style: fontstyle.Bold, Color: colorInk, Top: 9, Left: 4}),
				text.New(o.Supplier.Email, props.Text{Size: 8, Color: colorLink, Top: 15, Left: 4}),
			),
		),

		row.New(6),

		// Section label, matching the small caps labels on screen.
		row.New(6).Add(
			col.New(12).Add(
				text.New("ORDER ITEMS", props.Text{Size: 8, Style: fontstyle.Bold, Color: colorMuted, Top: 2}),
			),
		),

		// Table header: PRODUCT / QTY, underlined with a border-bottom rule.
		row.New(7).WithStyle(&props.Cell{
			BorderColor:     colorBorder,
			BorderType:      border.Bottom,
			BorderThickness: 0.3,
		}).Add(
			text.NewCol(8, "PRODUCT", props.Text{Size: 8, Style: fontstyle.Bold, Color: colorMuted, Top: 2, Left: 2}),
			text.NewCol(4, "QTY", props.Text{Size: 8, Style: fontstyle.Bold, Align: align.Right, Color: colorMuted, Top: 2}),
		),
	}

	// One row per product line, product names in the same blue used on
	// screen, divided by a thin bottom border like the on-screen list.
	for _, item := range o.Items {
		rows = append(rows, row.New(9).WithStyle(&props.Cell{
			BorderColor:     colorBorder,
			BorderType:      border.Bottom,
			BorderThickness: 0.2,
		}).Add(
			text.NewCol(8, item.ProductName, props.Text{Size: 10, Color: colorLink, Top: 2.5, Left: 2}),
			text.NewCol(4, strconv.Itoa(item.Qty), props.Text{Size: 10, Style: fontstyle.Bold, Align: align.Right, Color: colorInk, Top: 2.5}),
		))
	}

	rows = append(rows,
		row.New(9).Add(
			col.New(8).Add(
				text.New("Total quantity", props.Text{Size: 9, Color: colorMuted, Top: 3, Left: 2}),
			),
			col.New(4).Add(
				text.New(strconv.Itoa(o.TotalQuantity()), props.Text{Size: 12, Style: fontstyle.Bold, Align: align.Right, Color: colorInk, Top: 1.5}),
			),
		),
	)

	return rows
}

var (
	// colorBrand is the green used by the OrderDesk logo mark (#16A34A).
	colorBrand = &props.Color{Red: 22, Green: 163, Blue: 74}

	// colorAccent is the teal-green used for the "Send order" button and
	// the selected-supplier card border (#1D9E75).
	colorAccent = &props.Color{Red: 29, Green: 158, Blue: 117}

	// colorAccentSoft is the light mint background behind "Active" badges (#E1F5EE).
	colorAccentSoft = &props.Color{Red: 225, Green: 245, Blue: 238}

	// colorAccentDark is the dark teal text used inside "Active" badges (#0F6E5A).
	colorAccentDark = &props.Color{Red: 15, Green: 110, Blue: 90}

	// colorInk is the near-black used for headings and strong text (#1A1A18).
	colorInk = &props.Color{Red: 26, Green: 26, Blue: 24}

	// colorMuted is the soft gray used for uppercase labels and secondary text (#8E908A).
	colorMuted = &props.Color{Red: 142, Green: 144, Blue: 138}

	// colorLink is the blue used for product names and email addresses (#1A6DBC).
	colorLink = &props.Color{Red: 26, Green: 109, Blue: 188}

	// colorBorder is the light gray used for card borders and table dividers (#E4E4E2).
	colorBorder = &props.Color{Red: 228, Green: 228, Blue: 226}

	// colorWhite is plain white, used for text drawn on top of colorAccent.
	colorWhite = &props.Color{Red: 255, Green: 255, Blue: 255}
)
