package middleware

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

func JWTMiddleware(keyJWT string, next http.HandlerFunc) *http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "authorization error", http.StatusUnauthorized)
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			http.Error(w, "invalid token format", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			return []byte(keyJWT), nil
		})

		if err != nil {
			log.Error().Err(err).Msg("error parsing JWT token")
			http.Error(w, "Authorization prolbem. Please, try later", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			log.Error().Str("token", bearerToken[1][:5]).Msg("invalid JWT token")
			http.Error(w, "Authorization prolbem. Please, try later", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Error().Str("token", bearerToken[1][:5]).Msg("invalid JWT token payload")
			http.Error(w, "Authorization prolbem. Please, try later", http.StatusUnauthorized)
			return
		}

		userID, ok := claims["sub"].(float64)
		if !ok {
			log.Error().Msg("invalid userID")
			http.Error(w, "Authorization prolbem. Please, try later", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", int(userID))

		next.ServeHTTP(w, r.WithContext(ctx))
	}

}
