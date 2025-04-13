package auth_commands

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
