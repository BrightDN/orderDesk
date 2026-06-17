package invites

import (
	"errors"
	"fmt"
	"time"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/services/mailer"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func (is *InvitationService) SendCompany(c echo.Context) *errorHandling.AppError {

	company, appErr := is.companyService.Create(c)
	if appErr != nil {
		return appErr
	}

	token := ""
	var pqErr *pq.Error

	for i := 0; i < 5; i++ {
		var err error
		token, appErr = is.generateToken(32)
		if appErr != nil {
			return appErr
		}

		err = is.db.CreateInvite(c.Request().Context(), database.CreateInviteParams{
			Email:      company.Email,
			InviteType: string(Company),
			Token:      token,
			CompanyID:  company.ID,
			ExpiresAt:  time.Now().AddDate(0, 0, 2),
		})
		if err == nil {
			break
		}
		if i == 4 {
			return &errorHandling.AppError{
				Action:    "Sending company invitation - max attempts reached",
				LogError:  fmt.Errorf("Failed to generate unique token after 5 attempts: %v", err),
				UserError: errors.New("failed to send invitation"),
			}
		}

		if errors.As(err, &pqErr) &&
			pqErr.Code == "23505" &&
			pqErr.Constraint == "invites_token_key" {
			continue
		}
		return &errorHandling.AppError{
			Action:    "Sending company invitation - creating invite record",
			LogError:  err,
			UserError: errors.New("failed to send invitation"),
		}
	}

	content := is.getCompanyInvMail(token)

	mail := mailer.Mail{
		Receiver: company.Email,
		Subject:  fmt.Sprintf("Activate your %s account", is.identity.AppName),
		Body:     content,
	}

	if err := is.mailService.Send(mail); err != nil {
		return &errorHandling.AppError{
			Action:    "Sending company invitation email",
			LogError:  fmt.Errorf("Failed to send invitation mail: %v", err),
			UserError: errors.New("failed to send invitation email"),
		}
	}
	return nil
}
