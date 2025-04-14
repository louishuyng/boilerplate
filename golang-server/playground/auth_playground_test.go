package playground

import (
	"fmt"
	auth_commands "rz-server/internal/app/user/application/auth/commands"
	playground_setup "rz-server/playground/setup"
	"testing"

	"github.com/google/uuid"
)

func TestAuthPlayGround(t *testing.T) {
	service := playground_setup.MakeAuthService()

	_, _ = service.Register(auth_commands.RegisterUserCommand{
		Email:       "test2@test.com",
		Password:    "password",
		DisplayName: "test",
	})

	t.Log("Hello, playground!")
}

func TestAuthRefreshTokenPlayGround(t *testing.T) {
	service := playground_setup.MakeAuthService()

	authData, _ := service.Login(auth_commands.LoginUserCommand{
		Email:    "test2@test.com",
		Password: "password",
	})

	refreshData, _ := service.RefreshToken(auth_commands.RefreshTokenCommand{
		RefreshToken: authData.RefreshToken,
	})

	fmt.Printf("refreshData: %+v\n", refreshData)

	t.Log("Refresh token success")
}

func TestAuthLoginPlayGround(t *testing.T) {
	service := playground_setup.MakeAuthService()

	authData, _ := service.Login(auth_commands.LoginUserCommand{
		Email:    "test2@test.com",
		Password: "password",
	})

	fmt.Printf("authData: %+v\n", authData)

	t.Log("Login success")
}

func TestAuthLogoutPlayGround(t *testing.T) {
	service := playground_setup.MakeAuthService()

	err := service.Logout(auth_commands.LogoutUserCommand{
		UserID: uuid.MustParse("fd0eb9c2-72c8-4c10-b959-df79fe0946eb"),
	})

	if err != nil {
		t.Errorf("Logout failed: %v", err)
	}

	fmt.Printf("Logout success")
}
