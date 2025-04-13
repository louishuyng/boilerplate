package auth_api

import (
	"net/http"
	"rz-server/internal/app/user/application"
	auth_commands "rz-server/internal/app/user/application/auth/commands"
	"rz-server/internal/common/interfaces"
)

type AuthApi struct {
	service application.AuthService
	util    *interfaces.Util
	server  interfaces.Server
}

func New(
	server interfaces.Server,
	service application.AuthService,
	util *interfaces.Util,
) *AuthApi {
	u := new(AuthApi)
	u.service = service
	u.util = util
	u.server = server

	return u
}

func (u *AuthApi) Register() {
	u.server.POST("/user/auth/register", u.register)
	u.server.POST("/user/auth/login", u.login)
}

func (u *AuthApi) register(w http.ResponseWriter, r *http.Request) {
	u.util.Log.Info("Registering user", map[string]any{
		"method": r.Method,
	})

	command := auth_commands.RegisterUserCommand{
		Email:       "test@test.com",
		DisplayName: "Test User",
		Password:    "password",
	}

	_, _ = u.service.Register(command)

	w.Write([]byte("Example created"))
	w.WriteHeader(http.StatusCreated)
}

func (u *AuthApi) login(w http.ResponseWriter, r *http.Request) {
	u.util.Log.Info("Logging in user", map[string]any{
		"method": r.Method,
	})

	command := auth_commands.LoginUserCommand{
		Email:    "test@test.com",
		Password: "password",
	}

	_, _ = u.service.Login(command)

	w.Write([]byte("Logged in"))
	w.WriteHeader(http.StatusOK)
}
