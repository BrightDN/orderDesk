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

type allPermissions struct {
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
