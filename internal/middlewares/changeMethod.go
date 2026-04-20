package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ChangeMethod() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Method == http.MethodPost {
				switch method := c.FormValue("_method"); method {
				case http.MethodPut:
					fallthrough
				case http.MethodPatch:
					fallthrough
				case http.MethodDelete:
					c.Request().Method = method
				}
			}
			return next(c)
		}
	}
}
