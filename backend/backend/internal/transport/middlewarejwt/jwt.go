package middlewarejwt

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

type JWTMiddleware struct {
	keyJWT string
}

func NewJWTMiddleware(keyJWT string) *JWTMiddleware {
	return &JWTMiddleware{keyJWT: keyJWT}
}

func (j *JWTMiddleware) ParseJWTtoken(authHeader string) (jwt.MapClaims, error) {
	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
		return nil, errors.New("invalid token format")
	}

	token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
		return []byte(j.keyJWT), nil
	})

	if err != nil {
		log.Error().Err(err).Msg("error parsing JWT token")
	}

	if !token.Valid {
		log.Error().Str("token", bearerToken[1][:5]).Msg("invalid JWT token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Error().Str("token", bearerToken[1][:5]).Msg("invalid JWT token payload")
	}

	return claims, nil

}

func (j *JWTMiddleware) keyJWTMiddleware(isAdmin bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			var role = "user"
			var userID int

			if authHeader != "" {
				claims, err := j.ParseJWTtoken(authHeader)
				if err != nil {
					log.Error().Err(err).Msg("JWT token error")
					http.Error(w, "Authorization problem. Please, try later", http.StatusUnauthorized)
					return
				}

				role = "customer"
				if isAdmin {
					if r, ok := claims["role"].(string); ok && r == "admin" {
						role = "admin"
					} else {
						log.Error().Msg("User does not have admin permissions")
						http.Error(w, "You do not have permissions", http.StatusUnauthorized)
						return
					}
				} else {
					if r, ok := claims["role"].(string); ok {
						role = r
					}

				}

				var ok bool
				userID, ok = claims["sub"].(int)
				if !ok {
					log.Error().Msg("invalid userID")
					http.Error(w, "Authorization problem. Please try later", http.StatusUnauthorized)
					return
				}
			}
			ctx := context.WithValue(r.Context(), "userID", userID)
			if isAdmin {
				ctx = context.WithValue(ctx, "role", role)
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func (j *JWTMiddleware) JWTMiddlewareCustomer() func(http.Handler) http.Handler {
	return j.keyJWTMiddleware(false)
}

func (j *JWTMiddleware) JWTMiddlewareAdmin() func(http.Handler) http.Handler {
	return j.keyJWTMiddleware(true)
}
