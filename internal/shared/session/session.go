package session

import "errors"

var ErrNotFound = errors.New("value not found in session")
var ErrBadRetrieval = errors.New("failed to retrieve sessiondata")

type SessionData struct {
	UserID         int32
	RoleName       string
	CompanyID      int32
	IsMultiCompany bool
	IsSiteAdmin    bool
}

const (
	UserIDKey    = "UserID"
	RoleNameKey  = "RoleName"
	CompanyIDKey = "CompanyID"
	MultiCompKey = "IsMultiCompany"
	SiteAdminKey = "IsSiteAdmin"
)
