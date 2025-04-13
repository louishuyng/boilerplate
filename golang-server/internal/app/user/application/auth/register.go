package auth_service

import (
	"errors"
	"fmt"
	auth_commands "rz-server/internal/app/user/application/auth/commands"
	auth_data "rz-server/internal/app/user/application/auth/data"
	auth_store_data "rz-server/internal/app/user/infra/store/sql/auth/data"
	user_store_data "rz-server/internal/app/user/infra/store/sql/user/data"
)

func (s *AuthService) Register(command auth_commands.RegisterUserCommand) (*auth_data.AuthData, error) {
	user := s.userStore.GetUserByEmail(command.Email)

	fmt.Println("HH", user)

	if user != nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := s.auth.HashPassword(command.Password)

	if err != nil {
		return nil, err
	}

	newUser := s.userStore.CreateUser(user_store_data.CreateUserBody{
		Email:       command.Email,
		Password:    hashedPassword,
		DisplayName: command.DisplayName,
	})

	if newUser == nil {
		return nil, errors.New("failed to create user")
	}

	refreshToken, expiredAt, err := s.auth.GenerateRefreshToken(newUser.Id)

	if err != nil {
		return nil, err
	}

	s.authStore.SaveRefreshToken(auth_store_data.RefreshTokenBody{
		UserID:    newUser.Id,
		Token:     refreshToken,
		ExpiredAt: expiredAt,
	})

	accessToken, err := s.auth.GenerateAccessToken(newUser.Id)

	if err != nil {
		return nil, err
	}

	return &auth_data.AuthData{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil
}
