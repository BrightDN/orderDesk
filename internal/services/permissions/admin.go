package permissions

func (p *PermissionsService) getAdminPermissions() Permissions {
	return Permissions{
		CanViewOrders:       true,
		CanPlaceOrders:      true,
		CanViewSuppliers:    true,
		CanEditSuppliers:    true,
		CanViewCompany:      true,
		CanEditCompany:      true,
		CanViewOrderHistory: true,
	}
}
