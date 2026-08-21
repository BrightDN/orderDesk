package permissions

// GetCustomPermissions builds a permission set from the custom permissions
// returned with the employee role query.
func (p *PermissionsService) GetCustomPermissions(customPermissions []string) Permissions {
	permissions := Permissions{}

	for _, permission := range customPermissions {
		switch PermKey(permission) {
		case CanViewOrders:
			permissions.CanViewOrders = true
		case CanPlaceOrders:
			permissions.CanPlaceOrders = true
		case CanViewSuppliers:
			permissions.CanViewSuppliers = true
		case CanEditSuppliers:
			permissions.CanEditSuppliers = true
		case CanViewCompany:
			permissions.CanViewCompany = true
		case CanEditCompany:
			permissions.CanEditCompany = true
		case CanViewOrderHistory:
			permissions.CanViewOrderHistory = true
		}
	}

	return permissions
}
