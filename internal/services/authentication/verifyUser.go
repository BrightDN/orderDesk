package authentication

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/labstack/echo/v4"
)

func (auth *AuthenticationService) VerifyUser(c echo.Context, email, password string) (database.User, *errorHandling.AppError) {

	user, err := auth.queries.GetUserByMail(c.Request().Context(), email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return database.User{}, &errorHandling.AppError{
				Action:    "Verifying user credentials",
				LogError:  fmt.Errorf("User not found with email: %s", email),
				UserError: errors.New("invalid credentials"),
			}
		}
		return database.User{}, &errorHandling.AppError{
			Action:    "Fetching user by email",
			LogError:  err,
			UserError: errors.New("internal error"),
		}
	}

	isSame, appErr := auth.comparePasswordHash(password, user.Password)
	if appErr != nil {
		return database.User{}, appErr
	}
	if !isSame {
		return database.User{}, &errorHandling.AppError{
			Action:    "Verifying user credentials",
			LogError:  fmt.Errorf("Password mismatch for email: %s", email),
			UserError: errors.New("invalid credentials"),
		}
	}

	return user, nil
}
