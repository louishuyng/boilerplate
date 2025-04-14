package application_error

import (
	"fmt"
	"rz-server/internal/common/interfaces"
)

type Error struct {
	domain    string
	errorType ErrorType
	subDomain string
	code      int16
	message   string
	context   map[string]any
}

var _ interfaces.ApplicationError = (*Error)(nil)

func NewError(domain string, subDomain string, code int16, errorType ErrorType, message string, args ...any) *Error {
	return &Error{
		domain:    domain,
		errorType: errorType,
		subDomain: subDomain,
		code:      code,
		message:   fmt.Sprintf(message, args...),
		context:   make(map[string]any),
	}
}

func (e *Error) GetKey() string {
	return fmt.Sprintf("%s.%s.%d", e.domain, e.subDomain, e.code)
}

func (e *Error) GetMessage() string {
	return e.message
}

func (e *Error) GetContext() map[string]any {
	return e.context
}

func (e *Error) GetErrorType() string {
	return string(e.errorType)
}
