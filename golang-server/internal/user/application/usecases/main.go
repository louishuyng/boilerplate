package usecases

import (
	application "go-server/internal/user/application"
	domain "go-server/internal/user/domain"
	infra "go-server/internal/user/infra"
)

var _ application.UserService = (*UseCases)(nil)

type UseCases struct {
	store infra.UserStore
	core  domain.UserEntity
}

func NewUseCases(store infra.UserStore, core domain.UserEntity) *UseCases {
	return &UseCases{store, core}
}
