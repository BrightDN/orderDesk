package orders

import (
	"time"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func (os *OrderService) buildPDFHeader(order Order, t time.Time) []core.Row {

	return []core.Row{
		row.New(16).Add(
			image.NewFromFileCol(6, "assets/images/logo/orderdesk_logo.png", props.Rect{
				Center:  true,
				Percent: 80,
			}),

			col.New(5).Add(
				text.New("PURCHASE ORDER", props.Text{Size: 7, Style: fontstyle.Bold, Align: align.Right, Color: colorMuted, Top: 1}),
				text.New(order.ID, props.Text{Size: 11, Style: fontstyle.Bold, Align: align.Right, Color: colorInk, Top: 5.5}),
				text.New(t.Format("Jan 2, 2006"), props.Text{Size: 7, Align: align.Right, Color: colorMuted, Top: 11}),
			),
		),
		row.New(3),
		row.New(1).Add(
			line.NewCol(12, props.Line{Thickness: 0.7, Color: colorAccent, SizePercent: 100}),
		),
		row.New(4),
	}
}
