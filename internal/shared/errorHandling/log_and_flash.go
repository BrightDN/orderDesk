package errorHandling

import (
	"fmt"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/labstack/echo/v4"
)

func Log_and_flash(c echo.Context, err AppError) error {
	fmt.Printf("Logging type: error\nEncountered at: %s\nError: %v", err.Action, err.LogError)
	if err.UserError != nil {
		if flashErr := flash.Set(c, flash.Error, err.UserError.Error()); flashErr != nil {
			return flashErr
		}
	}
	return nil
}
