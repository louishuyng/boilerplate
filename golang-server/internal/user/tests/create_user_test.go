package integration_test

import (
	"go-server/internal/user/application"
	"go-server/internal/user/application/service"
	"go-server/internal/user/domain"
	"net/http"
	"testing"
)

func TestCreateUserAPI(t *testing.T) {
	t.Run("When creating user with valid data", func(t *testing.T) {

	})
}

func createAPI() {
	server := http.NewServeMux()
	// Store := store.TestContainerStore()
	core := domain.NewUserDomain()
	application := service.NewUserService(store, core)

	api := NewUserApi(server, createApplication())
}
