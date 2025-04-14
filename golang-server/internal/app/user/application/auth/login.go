package auth_service

import (
	"errors"
	"fmt"
	auth_commands "rz-server/internal/app/user/application/auth/commands"
	auth_data "rz-server/internal/app/user/application/auth/data"
	auth_store_data "rz-server/internal/app/user/infra/store/sql/auth/data"
)

func (s *AuthService) Login(command auth_commands.LoginUserCommand) (*auth_data.AuthData, error) {
	email := command.Email
	password := command.Password

	user := s.userStore.GetUserByEmail(email)

	fmt.Printf("user: %+v\n", user)

	if user == nil {
		return nil, errors.New("401: user not found")
	}

	isPasswordCorrect := s.auth.ComparePassword(password, user.Password)

	if !isPasswordCorrect {
		return nil, errors.New("401: invalid password")
	}

	existingRefreshToken := s.authStore.GetRefreshTokenByUserID(user.Id)

	refreshToken := ""

	if existingRefreshToken != nil {
		s.authStore.UpdateRefreshTokenExpiredAt(existingRefreshToken.ID, s.auth.GetExpiredAtAfter())
		refreshToken = existingRefreshToken.Token
	} else {
		newRefreshToken, expiredAt, err := s.auth.GenerateRefreshToken(user.Id)

		if err != nil {
			return nil, err
		}

		s.authStore.SaveRefreshToken(auth_store_data.RefreshTokenBody{
			UserID:    user.Id,
			Token:     newRefreshToken,
			ExpiredAt: expiredAt,
		})
		refreshToken = newRefreshToken
	}

	accessToken, err := s.auth.GenerateAccessToken(user.Id)

	if err != nil {
		return nil, err
	}

	return &auth_data.AuthData{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil
}
