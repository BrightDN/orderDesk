package errorHandling

import (
	"fmt"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/labstack/echo/v4"
)

func log_and_flash(c echo.Context, error AppError) error {
	fmt.Printf("Logging type: error\nEncountered at: %s\nError: %v", error.Action, error.LogError)
	if error.UserError != nil {
		if flashErr := flash.Set(c, flash.Error, error.UserError.Error()); flashErr != nil {
			return flashErr
		}
	}
	return nil
}
