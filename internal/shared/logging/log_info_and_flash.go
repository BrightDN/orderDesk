package logging

import (
	"fmt"
	"strings"
	"time"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/labstack/echo/v4"
)

func Log_info_and_flash(c echo.Context, logMsg, userMsg string) error {
	if strings.TrimSpace(logMsg) != "" {
		fmt.Printf("Logging type: info\nMessage: %s\nTimestamp: %v", logMsg, time.Now())
	}
	if strings.TrimSpace(userMsg) != "" {
		if flashErr := flash.Set(c, flash.Pass, userMsg); flashErr != nil {
			return flashErr
		}
	}
	return nil
}

func Log_info_and_flash_trigger(c echo.Context, logMsg, userMsg string) error {
	if strings.TrimSpace(logMsg) != "" {
		fmt.Printf("Logging type: info\nMessage: %s\nTimestamp: %v", logMsg, time.Now())
	}
	if strings.TrimSpace(userMsg) != "" {
		if flashErr := flash.Trigger(c, flash.Pass, userMsg); flashErr != nil {
			return flashErr
		}
	}
	return nil

}
