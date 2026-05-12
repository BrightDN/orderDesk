package invites

import "errors"

var ErrMaxAttempts = errors.New("maximum attempts passed")
var ErrInviteCreation = errors.New("failed to generate an invitation")
var ErrTokenCreation = errors.New("failed to generate a token")

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
