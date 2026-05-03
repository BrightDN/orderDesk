package configs

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	message := "Something went wrong on our end. Try again later."
	if code == http.StatusNotFound {
		message = "The page you're looking for doesn't exist."
	}

	c.Render(code, "error", map[string]any{
		"Code":    code,
		"Message": message,
	})
}
