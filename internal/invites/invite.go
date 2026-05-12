package invites

type Invite struct {
	IType       Type
	Url         string
	InviteeName string
	InviteeMail string
	ExpiryDate  string
	IsExpired   bool
	IsUsed      bool
	ID          int
}

type Type string

const (
	Employee Type = "employee"
	Company  Type = "company"
)
