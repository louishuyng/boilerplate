package interfaces

type ApplicationError interface {
	GetKey() string
	GetMessage() string
	GetContext() map[string]any
	GetErrorType() string
}

type ApplicationErrorManager interface {
	New(code int16, args ...any) ApplicationError

	RegisterBusinessError(code int16, message string, args ...any)
	RegisterValidationError(code int16, message string, args ...any)
	RegisterInvalidInputError(code int16, message string, args ...any)
	RegisterNotFoundError(code int16, message string, args ...any)
	RegisterNotAllowedError(code int16, message string, args ...any)
}
