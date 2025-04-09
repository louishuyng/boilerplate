package middlewares

import (
	"net/http"
	"rz-server/internal/common/interfaces"
)

func NewLoggingMiddleware(log interfaces.LogUtil) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		middleware := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Info("Request received", map[string]any{
				"method": r.Method,
				"url":    r.URL.String(),
			})
			next.ServeHTTP(w, r)
		})

		return middleware
	}
}
