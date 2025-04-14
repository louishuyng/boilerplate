package example_service

import (
	application "rz-server/internal/app/example/application"
	example_errors "rz-server/internal/app/example/application/example/errors"
	domain "rz-server/internal/app/example/domain"
	store "rz-server/internal/app/example/infra/store"
	"rz-server/internal/common/interfaces"
)

var _ application.ExampleService = (*ExampleService)(nil)

type ExampleService struct {
	store   store.ExampleStore
	example domain.Example
	errors  interfaces.ApplicationErrorManager
}

func New(store store.ExampleStore, example domain.Example) *ExampleService {
	application_error := example_errors.New()
	application_error.RegisterAllErrors()

	return &ExampleService{store, example, application_error}
}
