package invites

import "time"

type invite struct {
	IType       Type
	Url         string
	InviteeName string
	InviteeMail string
	ExpiryDate  time.Time
	IsExpired   bool
	ID          int
}

type Type string

const (
	Employee Type = "employee"
	Company  Type = "company"
)
