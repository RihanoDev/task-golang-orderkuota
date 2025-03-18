package middleware

import (
	"context"
	"net/http"
	"orderkuota/utils"
	"strings"
)

type contextKey string

const UserIDKey contextKey = "userID"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.ErrorResponse(w, http.StatusUnauthorized, "failed", "Authorization header is required")
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			utils.ErrorResponse(w, http.StatusUnauthorized, "failed", "Invalid token format")
			return
		}

		tokenString := tokenParts[1]

		userID, err := utils.VerifyToken(tokenString)
		if err != nil {
			utils.ErrorResponse(w, http.StatusUnauthorized, "failed", "Invalid or expired token")
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
