package service

import (
	application "rz-server/internal/app/example/application"
	domain "rz-server/internal/app/example/domain"
	store "rz-server/internal/app/example/infra/store"
)

var _ application.ExampleService = (*ExampleService)(nil)

type ExampleService struct {
	store store.ExampleStore
	core  domain.ExampleDomain
}

func NewExampleService(store store.ExampleStore, core domain.ExampleDomain) *ExampleService {
	return &ExampleService{store, core}
}
