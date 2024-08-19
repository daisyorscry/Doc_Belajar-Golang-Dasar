package auth

import (
	exception "RESTApi/Helper/Exception"
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
			exception.ServiceErr(fmt.Errorf("redirect"), "redirect to login", "unauthorized")
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			exception.ServiceErr(fmt.Errorf("redirect"), "redirect to login", "unauthorized")

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
			exception.ServiceErr(fmt.Errorf("redirect"), "redirect to login", "unauthorized")

			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Set user information in context for further use
			ctx := context.WithValue(r.Context(), "user", claims["username"])
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		} else {
			exception.ServiceErr(fmt.Errorf("redirect"), "redirect to login", "unauthorized")

			return
		}
	})
}
