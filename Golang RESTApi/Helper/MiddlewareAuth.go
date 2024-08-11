package helper

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("x-api-token")

		if tokenHeader == "" {
			http.Error(w, "Missing auth token", http.StatusUnauthorized)
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 || splitted[0] != "Bearer" {
			http.Error(w, "Invalid/Malformed auth token", http.StatusUnauthorized)
			return
		}

		tokenPart := splitted[1]

		token, err := jwt.Parse(tokenPart, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrNotSupported
			}
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid/Malformed auth token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, "Invalid/Malformed auth token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "username", claims["username"])
		ctx = context.WithValue(ctx, "userId", int(claims["userId"].(float64)))
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
