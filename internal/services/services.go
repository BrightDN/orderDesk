package services

import (
	"github.com/brightDN/orderDesk/internal/configs"
	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/services/authentication"
	"github.com/brightDN/orderDesk/internal/services/companies"
	"github.com/brightDN/orderDesk/internal/services/invites"
	"github.com/brightDN/orderDesk/internal/services/mailer"
)

type Services struct {
	Mailer      *mailer.MailerService
	Companies   *companies.CompanyService
	Invitations *invites.InvitationService
	Auth        *authentication.AuthenticationService
}

func NewServices(db *database.Queries, ms *mailer.MailerService, identiy *configs.IdentityConfig) *Services {
	companies := companies.NewCompanyService(db)
	invitations := invites.NewInvitationService(db, ms, companies, identiy)
	auth := authentication.NewAuthService(db)
	return &Services{
		Mailer:      ms,
		Companies:   companies,
		Invitations: invitations,
		Auth:        auth,
	}
}
