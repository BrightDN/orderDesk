package errorHandling

type AppError struct {
	Action    string
	LogError  error
	UserError error
}

// Error implements the error interface
func (ae *AppError) Error() string {
	if ae == nil {
		return "an error occurred"
	}

	if ae.UserError != nil {
		return ae.UserError.Error()
	}
	return "an error occurred"
}
