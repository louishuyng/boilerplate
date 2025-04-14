package auth_api

import (
	"encoding/json"
	"net/http"
	"rz-server/internal/app/user/application"
	auth_commands "rz-server/internal/app/user/application/auth/commands"
	"rz-server/internal/common/interfaces"

	"github.com/google/uuid"
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
	u.server.POST("/user/auth/logout", u.logout)
	u.server.POST("/user/auth/refresh-token", u.refreshToken)
}

func (u *AuthApi) register(w http.ResponseWriter, r *http.Request) {
	u.util.Log.Info("Registering user", map[string]any{
		"method": r.Method,
	})

	command := auth_commands.RegisterUserCommand{
		Email:       r.FormValue("email"),
		DisplayName: r.FormValue("display_name"),
		Password:    r.FormValue("password"),
	}

	_, _ = u.service.Register(command)

	w.Write([]byte("Example created"))
	w.WriteHeader(http.StatusCreated)
}

func (u *AuthApi) login(w http.ResponseWriter, r *http.Request) {
	u.util.Log.Info("Logging in user", map[string]any{
		"method": r.Method,
	})

	var command auth_commands.LoginUserCommand
	err := json.NewDecoder(r.Body).Decode(&command)
	if err != nil {
		u.util.Log.Error("Error decoding login request", map[string]any{
			"error": err.Error(),
		})
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	authData, err := u.service.Login(command)
	if err != nil {
		u.util.Log.Error("Error logging in user", map[string]any{
			"error": err.Error(),
		})
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	jsonData, err := json.Marshal(authData)
	if err != nil {
		u.util.Log.Error("Error marshaling auth data", map[string]any{
			"error": err.Error(),
		})
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (u *AuthApi) logout(w http.ResponseWriter, r *http.Request) {
	u.util.Log.Info("Logging out user", map[string]any{
		"method": r.Method,
	})

	command := auth_commands.LogoutUserCommand{
		UserID: uuid.MustParse(r.FormValue("user_id")),
	}

	err := u.service.Logout(command)

	if err != nil {
		u.util.Log.Error("Error logging out user", map[string]any{
			"error": err.Error(),
		})
	}

	w.Write([]byte("Logged out"))
	w.WriteHeader(http.StatusOK)
}

func (u *AuthApi) refreshToken(w http.ResponseWriter, r *http.Request) {
	u.util.Log.Info("Refreshing token", map[string]any{
		"method": r.Method,
	})

	command := auth_commands.RefreshTokenCommand{
		RefreshToken: r.FormValue("refresh_token"),
	}

	_, _ = u.service.RefreshToken(command)

	w.Write([]byte("Token refreshed"))
	w.WriteHeader(http.StatusOK)
}
