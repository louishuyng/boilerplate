package api

import (
	"go-server/internal/user/application"
	"net/http"
)

type UserApi struct {
	service application.UserService
}

func NewUserApi(
	router *http.ServeMux,
	service application.UserService,
) *UserApi {
	u := new(UserApi)

	u.service = service

	router.Handle("/users/", http.HandlerFunc(u.createUser))

	return u
}

func (u *UserApi) createUser(w http.ResponseWriter, r *http.Request) {
	// Call service
}
