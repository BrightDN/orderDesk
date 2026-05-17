package invites

import (
	"strconv"
	"strings"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/labstack/echo/v4"
)

func convertIDParam(c echo.Context) (int32, error) {
	sid := c.Param("id")
	if len(strings.TrimSpace(sid)) == 0 {
		if flashErr := flash.Set(c, flash.Error, ErrUnexpectedValue.Error()); flashErr != nil {
			return 0, flashErr
		}
		return 0, ErrUnexpectedValue
	}

	id, err := strconv.Atoi(sid)
	if err != nil {
		if flashErr := flash.Set(c, flash.Error, ErrUnexpectedValue.Error()); flashErr != nil {
			return 0, flashErr
		}
		return 0, ErrUnexpectedValue
	}
	return int32(id), nil
}
