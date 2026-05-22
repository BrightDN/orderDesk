package invites

import (
	"fmt"
	"strings"
)

func (is *InvitationService) getCompanyInvMail(token string) string {
	link := fmt.Sprintf("https://www.%s/auth/company/%s", strings.ToLower(is.identity.BaseURL), token)
	content := fmt.Sprintf("Hello,\n\nThank you for choosing %s.\nYour company account has been created successfully. To complete the setup process, please activate your account using the link below:\n%s\nPlease note that this activation link will expire in 48 hours.\nIf you did not request this account, you can safely ignore this email.\nBest regards,\nThe %s team",
		is.identity.AppName,
		link,
		is.identity.AppName)

	return content
}
