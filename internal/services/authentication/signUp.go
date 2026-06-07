package authentication

import (
	"errors"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/labstack/echo/v4"
)

var errCreatingUser = errors.New("error: failed to create user")
var errCreatingCompanyEmployee = errors.New("error: failed to create company employee")
var errHashingPassword = errors.New("error: failed to hash password")

func (auth *AuthenticationService) SignUp(c echo.Context, email, password, name string, invitation *database.Invite) (*database.CompanyUser, error) {
	tx, err := auth.db.BeginTx(c.Request().Context(), nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	queries := database.New(tx)

	hashedPassword, err := auth.hashPassword(password)
	if err != nil {
		return nil, errHashingPassword
	}

	user, err := queries.CreateUser(c.Request().Context(), database.CreateUserParams{
		Email:    email,
		Password: hashedPassword,
		Name:     name,
	})
	if err != nil {
		return nil, errCreatingUser
	}

	empl, err := queries.CreateCompanyEmployee(c.Request().Context(), database.CreateCompanyEmployeeParams{
		UserID:      user.ID,
		CompanyID:   invitation.CompanyID,
		RoleID:      invitation.RoleID,
		DisplayName: name,
	})
	if err != nil {
		return nil, errCreatingCompanyEmployee
	}

	if err = queries.SetInviteUsed(c.Request().Context(), invitation.ID); err != nil {
		return nil, err
	}

	return &empl, tx.Commit()
}
