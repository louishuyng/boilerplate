package errors

import (
	"fmt"
	"rz-server/internal/common/interfaces"
)

type ErrorType string

const (
	BusinessError   ErrorType = "BusinessError"
	ValidationError ErrorType = "ValidationError"
	InvalidInput    ErrorType = "InvalidInput"
	NotFound        ErrorType = "NotFound"
	NotAllowed      ErrorType = "NotAllowed"
)

type ApplicationError struct {
	domain         string
	subDomain      string
	registeredCode map[string]*Error
}

var _ interfaces.ApplicationError = (*ApplicationError)(nil)

func NewApplicationError(domain string, subDomain string) *ApplicationError {
	return &ApplicationError{
		domain:         domain,
		subDomain:      subDomain,
		registeredCode: make(map[string]*Error),
	}
}

func (e *ApplicationError) New(code int16, args ...any) interfaces.Error {
	stringCode := fmt.Sprintf("%d", code)
	if error, ok := e.registeredCode[stringCode]; ok {
		makeErrorContext(error, args...)

		return error
	} else {
		panic(fmt.Sprintf("Error code %d not registered in domain %s and subdomain %s", code, e.domain, e.subDomain))
	}
}

func (e *ApplicationError) RegisterBusinessError(code int16, message string, args ...any) {
	e.registerError(code, BusinessError, message, args...)
}

func (e *ApplicationError) RegisterValidationError(code int16, message string, args ...any) {
	e.registerError(code, ValidationError, message, args...)
}

func (e *ApplicationError) RegisterInvalidInputError(code int16, message string, args ...any) {
	e.registerError(code, InvalidInput, message, args...)
}

func (e *ApplicationError) RegisterNotFoundError(code int16, message string, args ...any) {
	e.registerError(code, NotFound, message, args...)
}

func (e *ApplicationError) RegisterNotAllowedError(code int16, message string, args ...any) {
	e.registerError(code, NotAllowed, message, args...)
}

func (e *ApplicationError) registerError(code int16, errorType ErrorType, message string, args ...any) {
	stringCode := fmt.Sprintf("%d", code)

	if _, ok := e.registeredCode[stringCode]; ok {
		panic(fmt.Sprintf("Error code %d registered in domain: %s and subdomain: %s", code, e.domain, e.subDomain))
	}

	error := NewError(e.domain, e.subDomain, code, errorType, message, args...)
	makeErrorContext(error, args...)

	e.registeredCode[stringCode] = error
}

func makeErrorContext(error *Error, args ...any) *Error {
	if len(args) > 1 {
		panic("Only one argument is allowed")
	}

	// Only try to use args[0] if args is not empty
	if len(args) > 0 {
		if contextMap, ok := args[0].(map[string]any); ok {
			error.context = contextMap
		}
	}

	return error
}
