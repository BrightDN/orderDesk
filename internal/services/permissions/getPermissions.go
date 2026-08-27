package permissions

import "strings"

func (ps *PermissionsService) GetPermissions(role string) Permissions {
	switch strings.ToLower(role) {
	case "superadmin":
	case "admin":
		return ps.getAdminPermissions()

	case "employee":
		return ps.getEmployeePermissions()
	case "custom":
		return Permissions{}
	}
	return Permissions{}
}
