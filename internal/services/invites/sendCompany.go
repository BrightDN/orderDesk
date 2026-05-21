package invites

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/services/mailer"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func (is *InvitationService) SendCompany(c echo.Context, org string) error {

	company, err := is.companyService.Create(c)
	if err != nil {
		return err
	}

	token := ""
	var pqErr *pq.Error

	for i := range 5 {
		token, err = is.generateToken(32)
		if err != nil {
			return err
		}

		if err := is.db.CreateInvite(c.Request().Context(), database.CreateInviteParams{
			Email:      company.Email,
			InviteType: string(Company),
			Token:      token,
			CompanyID:  company.ID,
			ExpiresAt:  time.Now().AddDate(0, 0, 2),
		}); err == nil {
			break
		}
		if i == 5 {
			return fmt.Errorf("%w: %v", ErrMaxAttempts, err)
		}

		if errors.As(err, &pqErr) &&
			pqErr.Code == "23505" &&
			pqErr.Constraint == "invites_token_key" {
			continue
		}
		return fmt.Errorf("%w: %v", ErrInviteCreation, err)
	}

	link := fmt.Sprintf("https://www.%s/auth/company/%s", strings.ToLower(is.identity.BaseURL), token)
	content := fmt.Sprintf("Hello,\n\nThank you for choosing %s.\nYour company account has been created successfully. To complete the setup process, please activate your account using the link below:\n%s\nPlease note that this activation link will expire in 48 hours.\nIf you did not request this account, you can safely ignore this email.\nBest regards,\nThe %s team",
		org,
		link,
		org)

	mail := mailer.Mail{
		Receiver: company.Email,
		Subject:  fmt.Sprintf("Activate your %s account", org),
		Body:     content,
	}

	if err := is.mailService.Send(mail); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("Could not send the invitationmail: %v", err))
	}
	return nil
}
