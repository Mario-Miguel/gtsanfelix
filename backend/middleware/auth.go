package middleware

import (
	"net/http"
	"strings"

	"gtsanfelix/backend/auth"
)

// RequireAuth is a middleware that validates the Bearer JWT token.
// Wrap any handler that needs authentication with this.
func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, `{"error":"no autorizado"}`, http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		if _, err := auth.ValidateToken(tokenStr); err != nil {
			http.Error(w, `{"error":"token inválido"}`, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
