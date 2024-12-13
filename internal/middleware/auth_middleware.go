package middleware

import (
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler, validToken string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "bearer ") {
			http.Error(w, "Unathorization: missing authorization tiken", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "bearer ")

		if token != validToken {
			http.Error(w, "Forbidden: invalid token", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
