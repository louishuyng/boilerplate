package store

import (
	auth_store_data "rz-server/internal/app/user/infra/store/sql/auth/data"
	user_store_data "rz-server/internal/app/user/infra/store/sql/user/data"
	"time"

	"github.com/google/uuid"
)

type AuthStore interface {
	DeleteRefreshTokenByUserID(userID uuid.UUID) error
	GetRefreshTokenByToken(token string) *auth_store_data.Data
	GetRefreshTokenByUserID(userID uuid.UUID) *auth_store_data.Data
	SaveRefreshToken(body auth_store_data.RefreshTokenBody) auth_store_data.Data
	UpdateRefreshTokenExpiredAt(id uuid.UUID, expiredAt time.Time) error
}

type UserStore interface {
	GetUserByEmail(email string) *user_store_data.Data
	CreateUser(body user_store_data.CreateUserBody) *user_store_data.Data
}
