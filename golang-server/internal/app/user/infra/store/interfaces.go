package store

import (
	auth_store_data "rz-server/internal/app/user/infra/store/sql/auth/data"
	user_store_data "rz-server/internal/app/user/infra/store/sql/user/data"
	"time"
)

type AuthStore interface {
	GetRefreshTokenByToken(token string) *auth_store_data.Data
	SaveRefreshToken(body auth_store_data.RefreshTokenBody) auth_store_data.Data
	UpdateRefreshTokenExpiredAt(token string, expiredAt time.Time) error
}

type UserStore interface {
	GetUserByEmail(email string) *user_store_data.Data
	CreateUser(body user_store_data.CreateUserBody) *user_store_data.Data
}
