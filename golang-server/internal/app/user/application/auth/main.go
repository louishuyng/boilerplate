package auth_service

import (
	application "rz-server/internal/app/user/application"
	domain "rz-server/internal/app/user/domain"
	store "rz-server/internal/app/user/infra/store"
)

type AuthService struct {
	authStore store.AuthStore
	userStore store.UserStore
	auth      domain.Auth
}

var _ application.AuthService = (*AuthService)(nil)

func New(authStore store.AuthStore, userStore store.UserStore, auth domain.Auth) *AuthService {
	return &AuthService{authStore, userStore, auth}
}
