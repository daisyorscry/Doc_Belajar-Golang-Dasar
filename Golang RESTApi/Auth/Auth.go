package auth

import (
	helper "RESTApi/Helper"
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("x-api-token")

		if tokenHeader == "" {
			helper.WriteJsonResponse(w, http.StatusUnauthorized, "UNAUTHORIZED", "")

			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 || splitted[0] != "Bearer" {
			helper.WriteJsonResponse(w, http.StatusUnauthorized, "UNAUTHORIZED", "")

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
			helper.WriteJsonResponse(w, http.StatusUnauthorized, "UNAUTHORIZED", "")

			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			helper.WriteJsonResponse(w, http.StatusUnauthorized, "UNAUTHORIZED", "")

			return
		}

		ctx := context.WithValue(r.Context(), "username", claims["username"])
		ctx = context.WithValue(ctx, "userId", int(claims["userId"].(float64)))
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
