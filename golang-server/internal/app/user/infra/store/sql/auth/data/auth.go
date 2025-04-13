package auth_store_data

import (
	"time"

	"github.com/google/uuid"
)

type Data struct {
	ID       uuid.UUID
	UserID   uuid.UUID
	Token    string
	ExpireAt time.Time
}

func New(id uuid.UUID, userID uuid.UUID, token string, expireAt time.Time) Data {
	return Data{
		ID:       id,
		UserID:   userID,
		Token:    token,
		ExpireAt: expireAt,
	}
}

type RefreshTokenBody struct {
	UserID    uuid.UUID
	Token     string
	ExpiredAt time.Time
}
