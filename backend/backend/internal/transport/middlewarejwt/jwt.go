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
			if authHeader == "" {
				http.Error(w, "authorization error", http.StatusUnauthorized)
				return
			}

			claims, err := j.ParseJWTtoken(authHeader)
			if err != nil {
				log.Error().Err(err).Msg("JWT token error")
				http.Error(w, "Authorization problem. Please, try later", http.StatusUnauthorized)
				return
			}

			if isAdmin {
				role, ok := claims["role"].(string)
				if !ok || role != "admin" {
					log.Error().Msg("User have not admin permissions")
					http.Error(w, "You does not have permissions to it", http.StatusUnauthorized)
					return
				}
			}

			userID, ok := claims["sub"].(float64)
			if !ok {
				log.Error().Msg("invalid userID")
				http.Error(w, "Authorization problem. Please try later", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "userID", int(userID))
			if isAdmin {
				ctx = context.WithValue(ctx, "role", claims["role"])
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func (j *JWTMiddleware) JWTMiddlewareUser() func(http.Handler) http.Handler {
	return j.keyJWTMiddleware(false)
}

func (j *JWTMiddleware) JWTMiddlewareAdmin() func(http.Handler) http.Handler {
	return j.keyJWTMiddleware(true)
}

//TODO: Остановился на полном рефакторинге JWT мидлвара и частичном добавлении функций в юзер
//TODO: ... репозитории, сервисы и иже с ними.
//TODO: Остается добить методы, чтобы те точно работали, навесить авторизацию разного уровня...
//TODO: ... дать одному юзеру админку, а второго оставить без и затестить руками. Далее написать...
//TODO: ... автотесты для всего этого дела и запушить в гит. После пойдут товары
