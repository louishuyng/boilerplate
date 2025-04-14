package example_errors

import (
	"rz-server/internal/common/errors/application_error"
	"rz-server/internal/common/interfaces"
)

type ExampleError struct {
	interfaces.ApplicationError
}

func New() *ExampleError {
	applicationError := application_error.New("example", "example")

	return &ExampleError{
		ApplicationError: applicationError,
	}
}
