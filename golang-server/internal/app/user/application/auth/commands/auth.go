package auth_commands

import "github.com/google/uuid"

type RegisterUserCommand struct {
	DisplayName string
	Email       string
	Password    string
}

type LoginUserCommand struct {
	Email    string
	Password string
}

type RefreshTokenCommand struct {
	RefreshToken string
}

type LogoutUserCommand struct {
	UserID uuid.UUID
}
