package configs

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := "Something went wrong on our end. Try again later."

	var he *echo.HTTPError
	if errors.As(err, &he) {
		code = he.Code
	}

	switch code {
	case http.StatusNotFound:
		message = "The page you're looking for doesn't exist."
	case http.StatusForbidden:
		message = "You don't have permission to access this page."
	case http.StatusUnauthorized:
		message = "Please log in to continue."
	case http.StatusBadRequest:
		message = "The request could not be processed."
	}

	// Actual error logging
	slog.Error(
		"HTTP request failed",
		"error", err,
		"code", code,
		"path", c.Request().URL.Path,
		"method", c.Request().Method,
	)

	renderErr := c.Render(code, "error", map[string]any{
		"Code":    code,
		"Message": message,
	})
	if renderErr != nil {
		c.Logger().Error(
			"failed to render error page",
			"error", renderErr,
			"template", "error",
			"code", code,
			"path", c.Request().URL.Path,
			"method", c.Request().Method,
		)

		if !c.Response().Committed {
			if err := c.String(code, fmt.Sprintf("Error %d: %s", code, message)); err != nil {
				c.Logger().Error(
					"failed to write fallback error response",
					"error", err,
				)
			}
		}
	}
}
