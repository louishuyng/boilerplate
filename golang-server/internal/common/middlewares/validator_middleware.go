package middlewares

import "net/http"

type Validator = func(r *http.Request) bool

func NewAuthorizationMiddleware(validator Validator) func(http.Handler) http.Handler {
	return newValidatorMiddleware(validator, http.StatusUnauthorized)
}

func newValidatorMiddleware(validator Validator, httpStatus int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		middleware := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !validator(r) {
				http.Error(w, "Invalid request", httpStatus)
				return
			}
			next.ServeHTTP(w, r)
		})

		return middleware
	}
}
