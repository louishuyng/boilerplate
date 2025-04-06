package service

import (
	application "go-server/internal/user/application"
	domain "go-server/internal/user/domain"
	store "go-server/internal/user/infra/store"
)

var _ application.UserService = (*UserService)(nil)

type UserService struct {
	store store.UserStore
	core  domain.UserDomain
}

func NewUserService(store store.UserStore, core domain.UserDomain) *UserService {
	return &UserService{store, core}
}
