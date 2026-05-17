package invites

import "errors"

var ErrMaxAttempts = errors.New("maximum attempts passed")
var ErrInviteCreation = errors.New("failed to generate an invitation")
var ErrTokenCreation = errors.New("failed to generate a token")
var ErrUnexpectedValue = errors.New("Unexpected value, action failed")
var ErrInternalError = errors.New("Something went wrong and we could not complete your request")
var ErrAlreadyAccepted = errors.New("The invite has already been accepted")

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
