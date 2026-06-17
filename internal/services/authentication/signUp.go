package authentication

import (
	"errors"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/labstack/echo/v4"
)

func (auth *AuthenticationService) SignUp(c echo.Context, email, password, name string, invitation *database.Invite) (*database.CompanyUser, *errorHandling.AppError) {
	tx, err := auth.db.BeginTx(c.Request().Context(), nil)
	if err != nil {
		return nil, &errorHandling.AppError{
			Action:    "Beginning database transaction for sign up",
			LogError:  err,
			UserError: errors.New("failed to process sign up"),
		}
	}
	defer tx.Rollback()

	queries := database.New(tx)

	hashedPassword, appErr := auth.hashPassword(password)
	if appErr != nil {
		return nil, appErr
	}

	user, err := queries.CreateUser(c.Request().Context(), database.CreateUserParams{
		Email:    email,
		Password: hashedPassword,
		Name:     name,
	})
	if err != nil {
		return nil, &errorHandling.AppError{
			Action:    "Creating user account during sign up",
			LogError:  err,
			UserError: errors.New("failed to create user"),
		}
	}

	empl, err := queries.CreateCompanyEmployee(c.Request().Context(), database.CreateCompanyEmployeeParams{
		UserID:      user.ID,
		CompanyID:   invitation.CompanyID,
		RoleID:      invitation.RoleID,
		DisplayName: name,
	})
	if err != nil {
		return nil, &errorHandling.AppError{
			Action:    "Creating company employee during sign up",
			LogError:  err,
			UserError: errors.New("failed to create company employee"),
		}
	}

	if err = queries.SetInviteUsed(c.Request().Context(), invitation.ID); err != nil {
		return nil, &errorHandling.AppError{
			Action:    "Marking invitation as used during sign up",
			LogError:  err,
			UserError: errors.New("failed to process invitation"),
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, &errorHandling.AppError{
			Action:    "Committing transaction for sign up",
			LogError:  err,
			UserError: errors.New("failed to complete sign up"),
		}
	}

	return &empl, nil
}
