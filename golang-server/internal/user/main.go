package user

import (
	"go-server/internal/user/api"
	"go-server/internal/user/application/service"
	"go-server/internal/user/domain"
	store "go-server/internal/user/infra"
	"net/http"
)

func RegisterUserServer(router *http.ServeMux) error {
	store := store.NewPostgresStore()
	domain := domain.NewUserDomain()

	service := service.NewUserService(store, domain)
	api.NewUserApi(router, service)

	return nil
}
