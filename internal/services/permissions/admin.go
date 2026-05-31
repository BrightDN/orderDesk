package permissions

func (p *PermissionsService) GetAdminPermissions() allPermissions {
	return allPermissions{
		CanViewOrders:       true,
		CanPlaceOrders:      true,
		CanViewSuppliers:    true,
		CanEditSuppliers:    true,
		CanViewCompany:      true,
		CanEditCompany:      true,
		CanViewOrderHistory: true,
	}
}
