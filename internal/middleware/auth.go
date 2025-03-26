package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/yunz-dev/crowdflare/internal/services"
)

// AuthMiddleware protects routes by verifying JWTs
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        claims, err := service.ValidateJWT(tokenString)
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        username, ok := claims["username"].(string)
        if !ok {
            http.Error(w, "Invalid token payload", http.StatusUnauthorized)
            return
        }

        // Store username in request context
        ctx := context.WithValue(r.Context(), "username", username)
        next.ServeHTTP(w, r.WithContext(ctx))
    }
}
