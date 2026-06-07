package authentication

import (
	"database/sql"
	"errors"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/labstack/echo/v4"
)

var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrInternalError = errors.New("internal error")

func (auth *AuthenticationService) VerifyUser(c echo.Context, email, password string) (database.User, error) {

	user, err := auth.queries.GetUserByMail(c.Request().Context(), email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return database.User{}, ErrInvalidCredentials
		}
		return database.User{}, ErrInternalError
	}

	isSame, err := auth.comparePasswordHash(password, user.Password)
	if err != nil {
		return database.User{}, ErrInternalError
	}
	if !isSame {
		return database.User{}, ErrInvalidCredentials
	}

	return user, nil
}
