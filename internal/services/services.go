package services

import (
	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/services/companies"
	"github.com/brightDN/orderDesk/internal/services/invites"
	"github.com/brightDN/orderDesk/internal/services/mailer"
)

type Services struct {
	Mailer      *mailer.MailerService
	Companies   *companies.CompanyService
	Invitations *invites.InvitationService
}

func NewServices(db *database.Queries, ms *mailer.MailerService) *Services {
	companies := companies.CompanyService{}
	invitations := invites.NewInvitationService(db, ms, &companies)
	return &Services{
		Mailer:      ms,
		Companies:   &companies,
		Invitations: invitations,
	}
}
