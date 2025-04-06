package api

import (
	"go-server/internal/user/application"
	"net/http"
	"net/http/httptest"
	"testing"
)

type FakeUserService struct {
	calls []application.CreateUserCommand
}

func (f *FakeUserService) CreateUser(command application.CreateUserCommand) (*application.User, error) {
	f.calls = append(f.calls, command)
	return &application.User{}, nil
}

func TestCreateUser(t *testing.T) {
	request, err := http.NewRequest("GET", "/users/", nil)
	if err != nil {
		t.Fatal(err)
	}

	service := &FakeUserService{}

	// Mock response
	response := httptest.NewRecorder()

	router := http.NewServeMux()

	server := NewUserApi(router, service)
	server.ServeHTTP(response, request)

	// Check the status code
	if response.Code != http.StatusCreated {
		t.Errorf("expected status code %d, got %d", http.StatusCreated, response.Code)
	}

	// Check the service call
	if len(service.calls) != 1 {
		t.Errorf("expected 1 call to CreateUser, got %d", len(service.calls))
	}
}
