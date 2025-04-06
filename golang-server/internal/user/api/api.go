package api

import (
	"go-server/internal/common"
	"go-server/internal/user/application"
	"net/http"
)

type UserApi struct {
	server  common.BasicServer
	service application.UserService
}

func NewUserApi(
	server common.BasicServer,
	service application.UserService,
) *UserApi {
	u := new(UserApi)

	u.service = service

	server.GET("/users", u.createUser)

	return u
}

func (u *UserApi) createUser(w http.ResponseWriter, r *http.Request) {
	command := application.CreateUserCommand{
		Name: "John",
	}

	_, _ = u.service.CreateUser(command)

	w.Write([]byte("User created"))
	w.WriteHeader(http.StatusCreated)
}
