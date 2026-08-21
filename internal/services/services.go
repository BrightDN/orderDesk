package services

import (
	"database/sql"

	"github.com/brightDN/orderDesk/internal/configs"
	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/services/authentication"
	"github.com/brightDN/orderDesk/internal/services/companies"
	"github.com/brightDN/orderDesk/internal/services/companies/orders"
	"github.com/brightDN/orderDesk/internal/services/companies/suppliers"
	"github.com/brightDN/orderDesk/internal/services/invites"
	"github.com/brightDN/orderDesk/internal/services/mailer"
	"github.com/brightDN/orderDesk/internal/services/permissions"
)

type Services struct {
	Mailer      *mailer.MailerService
	Companies   *companies.CompanyService
	Suppliers   *suppliers.SupplierService
	Invitations *invites.InvitationService
	Auth        *authentication.AuthenticationService
	Orders      *orders.OrderService
	Permissions *permissions.PermissionsService
}

func NewServices(queries *database.Queries, db *sql.DB, ms *mailer.MailerService, identiy *configs.IdentityConfig) *Services {
	companies := companies.NewCompanyService(queries)
	invitations := invites.NewInvitationService(queries, ms, companies, identiy)
	suppliers := suppliers.NewSupplierService(queries, db)
	auth := authentication.NewAuthService(queries, db)
	orders := orders.NewOrderService(queries, db)
	permissions := permissions.NewPermissionsService(queries)
	return &Services{
		Mailer:      ms,
		Companies:   companies,
		Suppliers:   suppliers,
		Invitations: invitations,
		Auth:        auth,
		Orders:      orders,
		Permissions: permissions,
	}
}
