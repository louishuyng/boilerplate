package auth_service

import (
	"errors"
	auth_commands "rz-server/internal/app/user/application/auth/commands"
)

func (s *AuthService) Logout(command auth_commands.LogoutUserCommand) error {
	userID := command.UserID

	existingRefreshToken := s.authStore.GetRefreshTokenByUserID(userID)

	if existingRefreshToken == nil {
		return errors.New("refresh token not found")
	}

	err := s.authStore.DeleteRefreshTokenByUserID(userID)

	if err != nil {
		return err
	}

	return nil
}
