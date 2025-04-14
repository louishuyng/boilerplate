package example_errors

import (
	"rz-server/internal/common/errors/application_error"
	"rz-server/internal/common/interfaces"
)

type ExampleError struct {
	interfaces.ApplicationErrorManager
}

func New() *ExampleError {
	applicationErrorManager := application_error.NewManager("example", "example")

	return &ExampleError{
		ApplicationErrorManager: applicationErrorManager,
	}
}
