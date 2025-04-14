package application_error

type ErrorType string

const (
	BusinessError   ErrorType = "BusinessError"
	ValidationError ErrorType = "ValidationError"
	InvalidInput    ErrorType = "InvalidInput"
	NotFound        ErrorType = "NotFound"
	NotAllowed      ErrorType = "NotAllowed"
	InternalError   ErrorType = "InternalError"
)
