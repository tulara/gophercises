package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/tulara/coffeeshop/auth"
)

func WithAuth(next func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Assume Bearer token, not technically required but common with e.g Auth0
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		username, err := auth.VerifyJWTToken(tokenString)
		if err != nil {
			fmt.Printf("Error verifying token: %v", err)
			http.Error(w, "JWT Token error", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "username", username)

		next(w, r.WithContext(ctx))
	}
}
