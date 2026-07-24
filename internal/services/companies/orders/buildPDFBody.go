package orders

import (
	"fmt"
	"strconv"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func (os *OrderService) buildPDFBody(o Order) []core.Row {
	rows := []core.Row{
		row.New(4),

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

		row.New(6).Add(
			col.New(12).Add(
				text.New("ORDER ITEMS", props.Text{Size: 8, Style: fontstyle.Bold, Color: colorMuted, Top: 2}),
			),
		),

		row.New(7).WithStyle(&props.Cell{
			BorderColor:     colorBorder,
			BorderType:      border.Bottom,
			BorderThickness: 0.3,
		}).Add(
			text.NewCol(8, "PRODUCT", props.Text{Size: 8, Style: fontstyle.Bold, Color: colorMuted, Top: 2, Left: 2}),
			text.NewCol(4, "QTY", props.Text{Size: 8, Style: fontstyle.Bold, Align: align.Right, Color: colorMuted, Top: 2}),
		),
	}

	for _, item := range o.Items {
		rows = append(rows, row.New(9).WithStyle(&props.Cell{
			BorderColor:     colorBorder,
			BorderType:      border.Bottom,
			BorderThickness: 0.2,
		}).Add(
			text.NewCol(8, item.ProductName, props.Text{Size: 10, Color: colorInk, Top: 2.5, Left: 2}),
			text.NewCol(4, strconv.Itoa(item.Qty), props.Text{Size: 10, Style: fontstyle.Bold, Align: align.Right, Color: colorInk, Top: 2.5}),
		))
	}

	rows = append(rows,
		row.New(9).Add(
			col.New(8).Add(
				text.New("Total quantity", props.Text{Size: 9, Color: colorMuted, Top: 3, Left: 2}),
			),
			col.New(4).Add(
				text.New(strconv.Itoa(o.totalQuantity()), props.Text{Size: 12, Style: fontstyle.Bold, Align: align.Right, Color: colorInk, Top: 1.5}),
			),
		),
	)

	return rows
}
