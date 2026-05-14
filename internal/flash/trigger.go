package flash

import (
	"encoding/json"
	"strings"

	"github.com/labstack/echo/v4"
)

func Trigger(c echo.Context, messageType MessageType, message string) error {
	triggers := map[string]any{}
	existing := strings.TrimSpace(c.Response().Header().Get(htmxTriggerHeader))
	if strings.HasPrefix(existing, "{") {
		_ = json.Unmarshal([]byte(existing), &triggers)
	} else if existing != "" {
		for eventName := range strings.SplitSeq(existing, ",") {
			if eventName = strings.TrimSpace(eventName); eventName != "" {
				triggers[eventName] = true
			}
		}
	}

	triggers["feedback"] = Flash{
		Type:    messageType,
		Message: message,
	}

	payload, err := json.Marshal(triggers)
	if err != nil {
		return err
	}

	c.Response().Header().Set(htmxTriggerHeader, string(payload))
	return nil
}
