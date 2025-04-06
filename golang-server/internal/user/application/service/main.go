package service

import (
	application "go-server/internal/user/application"
	domain "go-server/internal/user/domain"
	infra "go-server/internal/user/infra"
)

var _ application.UserService = (*UserService)(nil)

type UserService struct {
	store infra.UserStore
	core  domain.UserDomain
}

func NewUserService(store infra.UserStore, core domain.UserDomain) *UserService {
	return &UserService{store, core}
}
