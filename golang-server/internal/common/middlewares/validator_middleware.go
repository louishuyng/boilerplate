package middlewares

import (
	"context"
	"net/http"

	jwt "github.com/golang-jwt/jwt/v4"
)

type Validator = func(r *http.Request) bool

func NewJWTAuthorizationMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		middleware := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bearerToken := r.Header.Get("Authorization")

			if bearerToken == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			}

			bearerToken = bearerToken[len("Bearer "):]

			// Parse the JWT token
			_, err := jwt.Parse(bearerToken, func(token *jwt.Token) (any, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, http.ErrNotSupported
				}

				return []byte("your-256-bit-secret"), nil // Replace with your secret key
			})

			ctx := context.WithValue(r.Context(), "user", "123")

			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})

		return middleware
	}
}
