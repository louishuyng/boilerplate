package auth_service

import (
	"errors"
	auth_commands "rz-server/internal/app/user/application/auth/commands"
	auth_data "rz-server/internal/app/user/application/auth/data"
)

func (s *AuthService) RefreshToken(command auth_commands.RefreshTokenCommand) (*auth_data.AuthData, error) {
	refreshToken := command.RefreshToken

	refreshTokenData := s.authStore.GetRefreshTokenByToken(refreshToken)

	if refreshTokenData == nil {
		return nil, errors.New("refresh token not found")
	}

	err := s.auth.ValidateExpired(refreshTokenData.ExpireAt)

	if err != nil {
		return nil, err
	}

	err = s.authStore.UpdateRefreshTokenExpiredAt(refreshTokenData.ID, s.auth.GetExpiredAtAfter())

	if err != nil {
		return nil, err
	}

	accessToken, err := s.auth.GenerateAccessToken(refreshTokenData.UserID)

	if err != nil {
		return nil, err
	}

	return &auth_data.AuthData{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
