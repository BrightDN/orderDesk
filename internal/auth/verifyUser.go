package auth

import (
	"net/http"
	"net/mail"

	"github.com/brightDN/orderDesk/internal/configs"
	"github.com/brightDN/orderDesk/internal/database"
	"github.com/labstack/echo/v4"
)

func verifyUser(c echo.Context, cfg *configs.Config) error {
	email := c.FormValue("email")
	password := c.Request().PostFormValue("password")

	if email == "" {
		return c.Render(http.StatusBadRequest, "login", map[string]any{
			"error": "Email is required",
		})
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return c.Render(http.StatusBadRequest, "login", map[string]any{
			"error": "Invalid email address",
		})
	}

	if password == "" {
		return c.Render(http.StatusBadRequest, "login", map[string]any{
			"error": "Password is required",
		})
	}

	user, err := cfg.Db.GetUserByMail(c.Request().Context(), email)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "login", map[string]any{
			"error": "Something went wrong, please try again later",
		})
	}
	if (user == database.User{}) {
		return c.Render(http.StatusUnauthorized, "login", map[string]any{
			"error": "Invalid credentials",
		})
	}

	isSame, err := ComparePasswordHash(password, user.Password)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "login", map[string]any{
			"error": "Something went wrong, please try again later",
		})
	}
	if !isSame {
		return c.Render(http.StatusUnauthorized, "login", map[string]any{
			"error": "Invalid credentials",
		})
	}

	return c.Render(http.StatusOK, "suppliers", nil)
}
