package permissions

import "github.com/brightDN/orderDesk/internal/database"

type PermissionsService struct {
	db *database.Queries
}

func NewPermissionsService(db *database.Queries) *PermissionsService {
	return &PermissionsService{
		db: db,
	}
}

// Permissions contains the actions the current employee is allowed to perform.
// It is stored in the Echo context by LoadPermissions.
type Permissions struct {
	// order page
	CanViewOrders  bool
	CanPlaceOrders bool

	// suppliers page
	CanViewSuppliers bool
	CanEditSuppliers bool

	// company page
	CanViewCompany bool
	CanEditCompany bool

	// order history page
	CanViewOrderHistory bool
}

const ContextKey = "permissions"

type PermKey string

const (
	CanViewOrders       PermKey = "canViewOrders"
	CanPlaceOrders      PermKey = "canPlaceOrders"
	CanViewSuppliers    PermKey = "canViewSuppliers"
	CanEditSuppliers    PermKey = "canEditSuppliers"
	CanViewCompany      PermKey = "canViewCompany"
	CanEditCompany      PermKey = "canEditCompany"
	CanViewOrderHistory PermKey = "canViewOrderHistory"
)

func (s *PermissionsService) getPermKeys() []PermKey {
	return []PermKey{
		CanViewOrders,
		CanPlaceOrders,
		CanViewSuppliers,
		CanEditSuppliers,
		CanViewCompany,
		CanEditCompany,
		CanViewOrderHistory,
	}
}
