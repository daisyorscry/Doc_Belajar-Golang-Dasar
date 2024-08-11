package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")

		// Token is usually sent as "Bearer <token>", so we split it
		if tokenHeader == "" {
			http.Error(w, "Missing auth token", http.StatusUnauthorized)
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			http.Error(w, "Invalid/Malformed auth token", http.StatusUnauthorized)
			return
		}

		tokenPart := splitted[1] // The actual token

		token, err := jwt.Parse(tokenPart, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("Kdfg;dllkNOUB90-OKMON[3UHO'PMKJyhunsdsko;niu8093p'sl)(*&^BHn)"), nil
		})

		if err != nil {
			http.Error(w, "Invalid/Malformed auth token", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Set user information in context for further use
			ctx := context.WithValue(r.Context(), "user", claims["username"])
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Invalid/Malformed auth token", http.StatusUnauthorized)
			return
		}
	})
}
