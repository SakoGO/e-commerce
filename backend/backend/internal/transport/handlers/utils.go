package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

type Validator interface {
	ValidateStruct(interface{}) error
}

type JWTMiddleware interface {
	JWTMiddlewareUser() func(http.Handler) http.Handler
	JWTMiddlewareAdmin() func(http.Handler) http.Handler
}

func GetID(w http.ResponseWriter, r *http.Request) (int, error) {
	idStr := chi.URLParam(r, "id")
	shopID, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error().Err(err).Msg("Invalid id format")
		http.Error(w, "invalid id format", http.StatusBadRequest)
		return 0, err
	}
	return shopID, nil
}
