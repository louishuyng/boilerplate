package application

import (
	auth_commands "rz-server/internal/app/user/application/auth/commands"
	auth_data "rz-server/internal/app/user/application/auth/data"
)

type AuthService interface {
	Register(command auth_commands.RegisterUserCommand) (*auth_data.AuthData, error)
	Login(command auth_commands.LoginUserCommand) (*auth_data.AuthData, error)
	RefreshToken(command auth_commands.RefreshTokenCommand) (*auth_data.AuthData, error)
	Logout(command auth_commands.LogoutUserCommand) error
}
