package permissions

func (p *PermissionsService) GetEmployeePermissions() Permissions {
	return Permissions{
		CanViewOrders:       true,
		CanPlaceOrders:      true,
		CanViewSuppliers:    true,
		CanEditSuppliers:    false,
		CanViewCompany:      false,
		CanEditCompany:      false,
		CanViewOrderHistory: true,
	}
}
