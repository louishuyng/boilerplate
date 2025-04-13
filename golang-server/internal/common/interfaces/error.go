package interfaces

type Error interface {
	GetKey() string
	GetMessage() string
	GetContext() map[string]any
	GetErrorType() string
}

type ApplicationError interface {
	New(code int16, args ...any) Error

	RegisterBusinessError(code int16, message string, args ...any)
	RegisterValidationError(code int16, message string, args ...any)
	RegisterInvalidInputError(code int16, message string, args ...any)
	RegisterNotFoundError(code int16, message string, args ...any)
	RegisterNotAllowedError(code int16, message string, args ...any)
}
