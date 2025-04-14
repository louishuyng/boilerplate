package application_error

import (
	"fmt"
	"maps"
	"rz-server/internal/common/interfaces"
)

type Manager struct {
	domain         string
	subDomain      string
	registeredCode map[string]*Error
}

var _ interfaces.ApplicationErrorManager = (*Manager)(nil)

func NewManager(domain string, subDomain string) *Manager {
	registeredCode := make(map[string]*Error)

	maps.Copy(registeredCode, DEFAULT_ERRORS)

	return &Manager{
		domain:         domain,
		subDomain:      subDomain,
		registeredCode: registeredCode,
	}
}

func (e *Manager) New(code int16, args ...any) interfaces.ApplicationError {
	stringCode := fmt.Sprintf("%d", code)
	if err, ok := e.registeredCode[stringCode]; ok {
		makeErrorContext(err, args...)

		return err
	} else {
		panic(fmt.Sprintf("Error code %d not registered in domain %s and subdomain %s", code, e.domain, e.subDomain))
	}
}

func (e *Manager) RegisterBusinessError(code int16, message string, args ...any) {
	e.registerError(code, BusinessError, message, args...)
}

func (e *Manager) RegisterValidationError(code int16, message string, args ...any) {
	e.registerError(code, ValidationError, message, args...)
}

func (e *Manager) RegisterInvalidInputError(code int16, message string, args ...any) {
	e.registerError(code, InvalidInput, message, args...)
}

func (e *Manager) RegisterNotFoundError(code int16, message string, args ...any) {
	e.registerError(code, NotFound, message, args...)
}

func (e *Manager) RegisterNotAllowedError(code int16, message string, args ...any) {
	e.registerError(code, NotAllowed, message, args...)
}

func (e *Manager) registerError(code int16, errorType ErrorType, message string, args ...any) {
	stringCode := fmt.Sprintf("%d", code)

	if _, ok := e.registeredCode[stringCode]; ok {
		panic(fmt.Sprintf("Error code %d registered in domain: %s and subdomain: %s", code, e.domain, e.subDomain))
	}

	err := NewError(e.domain, e.subDomain, code, errorType, message, args...)
	makeErrorContext(err, args...)

	e.registeredCode[stringCode] = err
}

func makeErrorContext(err *Error, args ...any) *Error {
	if len(args) > 1 {
		panic("Only one argument is allowed")
	}

	// Only try to use args[0] if args is not empty
	if len(args) > 0 {
		if contextMap, ok := args[0].(map[string]any); ok {
			err.context = contextMap
		}
	}

	return err
}
