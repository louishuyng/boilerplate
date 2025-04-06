package api

import (
	"go-server/internal/user/application"
	"net/http"
)

type UserApi struct {
	service application.UserService
	http.Handler
}

func NewUserApi(
	router *http.ServeMux,
	service application.UserService,
) *UserApi {
	u := new(UserApi)

	u.service = service
	u.Handler = router

	router.Handle("/users/", http.HandlerFunc(u.createUser))

	return u
}

func (u *UserApi) createUser(w http.ResponseWriter, r *http.Request) {
	command := application.CreateUserCommand{
		Name: "John",
	}

	_, _ = u.service.CreateUser(command)

	w.WriteHeader(http.StatusCreated)
}
