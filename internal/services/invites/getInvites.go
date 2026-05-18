package invites

import (
	"fmt"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func (is *InvitationService) GetCompanyInvites(c echo.Context) []Invite {
	appName := "orderdesk"
	cinvs, _ := is.db.GetCompanyInvites(c.Request().Context())
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
