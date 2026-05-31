package permissions

func (p *PermissionsService) GetEmployeePermissions() allPermissions {
	return allPermissions{
		CanViewOrders:       true,
		CanPlaceOrders:      true,
		CanViewSuppliers:    true,
		CanEditSuppliers:    false,
		CanViewCompany:      false,
		CanEditCompany:      false,
		CanViewOrderHistory: true,
	}
}
