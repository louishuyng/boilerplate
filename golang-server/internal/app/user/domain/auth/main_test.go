package auth

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGenerateAccessToken(t *testing.T) {
	// Setup
	auth := New()
	userID := uuid.New()

	// Execute
	token, err := auth.GenerateAccessToken(userID)

	t.Run("should generate a valid access token", func(t *testing.T) {
		// Verify
		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		// Parse and validate the token
		parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		assert.NoError(t, err)
		assert.True(t, parsedToken.Valid)
	})

	t.Run("should have valid expiration time with in an hour", func(t *testing.T) {
		parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		assert.NoError(t, err)

		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		assert.True(t, ok)

		expTime, ok := claims["exp"].(float64)
		assert.True(t, ok)

		expirationTime := time.Unix(int64(expTime), 0)

		fmt.Println(expirationTime)

		// Check that expiration time is within an hour from now
		// The token is set to expire in 1 minute in the implementation
		// So we need to check that it's between now and now+~1 minute

		// Expire 22:00
		// Now: 21:00
		// Now after 1 hour >= Expire
		// Now after 30 minutes < Expire
		tolerateTime := time.Second * 5
		assert.True(t, time.Now().Add(time.Hour*1+tolerateTime).After(expirationTime), "Now after 1 hour should be after expiration time")
		assert.True(t, time.Now().Add(time.Minute*30).Before(expirationTime), "Now after 30 minutes should be before expiration time")
	})

}

func TestGenerateRefreshToken(t *testing.T) {
	auth := New()
	userID := uuid.New()

	token, _, err := auth.GenerateRefreshToken(userID)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	t.Run("should generate a valid refresh token", func(t *testing.T) {
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})

	t.Run("ID must be encrypted", func(t *testing.T) {
		assert.NotEqual(t, userID, token)
	})
}

func TestHashPassword(t *testing.T) {
	auth := New()
	password := "password"

	hashedPassword, err := auth.HashPassword(password)
	assert.NoError(t, err)

	t.Run("should generate a valid hashed password", func(t *testing.T) {
		assert.NoError(t, err)
		assert.NotEmpty(t, hashedPassword)
	})
}

func TestComparePassword(t *testing.T) {
	auth := New()
	password := "password"

	hashedPassword, err := auth.HashPassword(password)
	assert.NoError(t, err)

	t.Run("should return true if the password is correct", func(t *testing.T) {
		assert.True(t, auth.ComparePassword(password, hashedPassword))
	})

	t.Run("should return false if the password is incorrect", func(t *testing.T) {
		assert.False(t, auth.ComparePassword("wrong_password", hashedPassword))
	})
}
