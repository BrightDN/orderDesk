package flash

const (
	feedbackKey       = "feedback"
	separator         = "\x00"
	htmxTriggerHeader = "HX-Trigger"
)

type MessageType string

const (
	Error MessageType = "error"
	Pass  MessageType = "pass"
)

type Flash struct {
	Type    MessageType `json:"type"`
	Message string      `json:"message"`
}
