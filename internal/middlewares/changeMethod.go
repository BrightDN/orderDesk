package middlewares

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func changeMethod() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Method == http.MethodPost {
				method := strings.ToUpper(c.FormValue("_method"))

				switch method {
				case http.MethodPut, http.MethodPatch, http.MethodDelete:
					c.Request().Method = method
				}
			}
			return next(c)
		}
	}
}
