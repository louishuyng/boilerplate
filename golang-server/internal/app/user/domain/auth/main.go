package auth

import (
	"errors"
	domain "rz-server/internal/app/user/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("access-token-secret-key")

const REFRESH_TOKEN_EXPIRE_AT_AFTER = time.Hour * 24 * 30

type Entity struct {
}

var _ domain.Auth = (*Entity)(nil)

func New() *Entity {
	return &Entity{}
}

func (e *Entity) GenerateAccessToken(id uuid.UUID) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := claims.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (e *Entity) GenerateRefreshToken(id uuid.UUID) (string, time.Time, error) {
	byteID := []byte(id.String())

	token, err := bcrypt.GenerateFromPassword(byteID, bcrypt.DefaultCost)

	if err != nil {
		return "", time.Time{}, err
	}

	return string(token), time.Now().Add(REFRESH_TOKEN_EXPIRE_AT_AFTER), nil
}

func (e *Entity) HashPassword(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (e *Entity) ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (e *Entity) ValidateExpired(expiredAt time.Time) (time.Time, error) {
	if expiredAt.Before(time.Now()) {
		return time.Time{}, errors.New("token expired")
	}

	return time.Now().Add(REFRESH_TOKEN_EXPIRE_AT_AFTER), nil
}
