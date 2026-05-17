package companies

import (
	"errors"
	"strconv"
	"strings"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/labstack/echo/v4"
)

var ErrUnexpectedValue = errors.New("Unexpected value, action failed")
var ErrInternalError = errors.New("Something went wrong and we could not complete your request")

func convertIDParam(c echo.Context) (int32, error) {
	sid := c.Param("id")
	if len(strings.TrimSpace(sid)) == 0 {
		if flashErr := flash.Set(c, flash.Error, ErrUnexpectedValue.Error()); flashErr != nil {
			return 0, flashErr
		}
		return 0, ErrUnexpectedValue
	}

	id, err := strconv.ParseInt(sid, 10, 32)
	if err != nil {
		if flashErr := flash.Set(c, flash.Error, ErrUnexpectedValue.Error()); flashErr != nil {
			return 0, flashErr
		}
		return 0, ErrUnexpectedValue
	}
	return int32(id), nil
}
