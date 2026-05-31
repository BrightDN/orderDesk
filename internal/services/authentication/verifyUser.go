package authentication

import (
	"errors"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/labstack/echo/v4"
)

var errInvalidCredentials = errors.New("invalid credentials")
var errInternalError = errors.New("internal error")

func (auth *AuthenticationService) VerifyUser(c echo.Context, email, password string) (database.User, error) {

	user, err := auth.queries.GetUserByMail(c.Request().Context(), email)
	if err != nil {
		return database.User{}, errInternalError
	}
	if (user == database.User{}) {
		return database.User{}, errInvalidCredentials
	}

	isSame, err := auth.comparePasswordHash(password, user.Password)
	if err != nil {
		return database.User{}, errInternalError
	}
	if !isSame {
		return database.User{}, errInvalidCredentials
	}

	return user, nil
}
