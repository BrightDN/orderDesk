package invites

import (
	"fmt"
	"strings"
	"time"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/labstack/echo/v4"
)

func GetCompanyInvites(db *database.Queries, c echo.Context) []Invite {
	appName := "orderdesk"
	cinvs, _ := db.GetCompanyInvites(c.Request().Context())
	var invs = []Invite{}
	now := time.Now()
	for _, cinv := range cinvs {
		invs = append(invs, Invite{
			IType:       Type(cinv.InviteType),
			Url:         fmt.Sprintf("https://www.%s/invites/%s", strings.ToLower(appName), cinv.Token),
			InviteeName: cinv.CompanyName,
			InviteeMail: cinv.Email,
			ExpiryDate:  cinv.ExpiresAt.Format("02-01-2006"),
			IsExpired:   cinv.ExpiresAt.Before(now),
			IsUsed:      cinv.UsedAt.Valid,
			ID:          int(cinv.ID),
		})
	}

	return invs
}
