package orders

import (
	"fmt"
	"time"
)

func (os *OrderService) generateOrderID(now time.Time, id int32) string {
	return fmt.Sprintf("ORD-%s-%d", now.Format("20060102-150405"), id)
}
