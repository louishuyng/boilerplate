package domain

import (
	"time"

	"github.com/google/uuid"
)

type Auth interface {
	GenerateAccessToken(id uuid.UUID) (string, error)
	GenerateRefreshToken(id uuid.UUID) (string, time.Time, error)
	HashPassword(password string) (string, error)
	ComparePassword(password, hash string) bool
	ValidateExpired(expiredAt time.Time) (time.Time, error)
}
