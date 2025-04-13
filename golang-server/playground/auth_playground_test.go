package playground

import (
	"fmt"
	auth_commands "rz-server/internal/app/user/application/auth/commands"
	playground_setup "rz-server/playground/setup"
	"testing"
)

func TestAuthPlayGround(t *testing.T) {
	service := playground_setup.MakeAuthService()

	_, _ = service.Register(auth_commands.RegisterUserCommand{
		Email:       "test1@test.com",
		Password:    "password",
		DisplayName: "test",
	})

	t.Log("Hello, playground!")
}

func TestAuthLoginPlayGround(t *testing.T) {
	service := playground_setup.MakeAuthService()

	authData, _ := service.Login(auth_commands.LoginUserCommand{
		Email:    "test1@test.com",
		Password: "password",
	})

	fmt.Printf("authData: %+v\n", authData)

	t.Log("Login success")
}

func TestAuthRefreshTokenPlayGround(t *testing.T) {
	service := playground_setup.MakeAuthService()

	authData, _ := service.Login(auth_commands.LoginUserCommand{
		Email:    "test1@test.com",
		Password: "password",
	})

	authData, _ = service.RefreshToken(auth_commands.RefreshTokenCommand{
		RefreshToken: authData.RefreshToken,
	})

	fmt.Printf("authData: %+v\n", authData)

	t.Log("Refresh token success")
}
