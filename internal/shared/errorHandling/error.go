package errorHandling

type AppError struct {
	Action    string
	LogError  error
	UserError error
}
