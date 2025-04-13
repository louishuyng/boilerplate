package auth_service

import (
	"errors"
	auth_commands "rz-server/internal/app/user/application/auth/commands"
	auth_data "rz-server/internal/app/user/application/auth/data"
	auth_store_data "rz-server/internal/app/user/infra/store/sql/auth/data"
)

func (s *AuthService) Login(command auth_commands.LoginUserCommand) (*auth_data.AuthData, error) {
	email := command.Email
	password := command.Password

	user := s.userStore.GetUserByEmail(email)

	if user == nil {
		return nil, errors.New("user not found")
	}

	isPasswordCorrect := s.auth.ComparePassword(password, user.Password)

	if !isPasswordCorrect {
		return nil, errors.New("401: invalid password")
	}

	newRefreshToken, expiredAt, err := s.auth.GenerateRefreshToken(user.Id)

	if err != nil {
		return nil, err
	}

	s.authStore.SaveRefreshToken(auth_store_data.RefreshTokenBody{
		UserID:    user.Id,
		Token:     newRefreshToken,
		ExpiredAt: expiredAt,
	})

	accessToken, err := s.auth.GenerateAccessToken(user.Id)

	if err != nil {
		return nil, err
	}

	return &auth_data.AuthData{
		RefreshToken: newRefreshToken,
		AccessToken:  accessToken,
	}, nil
}
