package invites

import (
	"github.com/brightDN/orderDesk/internal/configs"
	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/services/companies"
	"github.com/brightDN/orderDesk/internal/services/mailer"
)

type InvitationService struct {
	mailService    *mailer.MailerService
	companyService *companies.CompanyService
	identity       *configs.IdentityConfig
	db             *database.Queries
}

func NewInvitationService(db *database.Queries, ms *mailer.MailerService, cs *companies.CompanyService, ic *configs.IdentityConfig) *InvitationService {
	return &InvitationService{
		mailService:    ms,
		db:             db,
		identity:       ic,
		companyService: cs,
	}
}
