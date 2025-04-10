package example_service

import (
	application "rz-server/internal/app/example/application"
	domain "rz-server/internal/app/example/domain"
	store "rz-server/internal/app/example/infra/store"
)

var _ application.ExampleService = (*ExampleService)(nil)

type ExampleService struct {
	store   store.ExampleStore
	example domain.Example
}

func New(store store.ExampleStore, example domain.Example) *ExampleService {
	return &ExampleService{store, example}
}
