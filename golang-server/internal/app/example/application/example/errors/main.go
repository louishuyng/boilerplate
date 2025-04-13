package example_errors

import (
	"rz-server/internal/common/errors"
	"rz-server/internal/common/interfaces"
)

type ExampleError struct {
	interfaces.ApplicationError
}

func New() *ExampleError {
	applicationError := errors.NewApplicationError("example", "example")

	return &ExampleError{
		ApplicationError: applicationError,
	}
}
