package handlers

import (
	"e-commerce/backend/internal/model"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

type AuthService interface {
	SignUP(user *model.User) error
	SignIN(email, password string) (string, error)
	FindByEmail(email string) (*model.User, error)
}

type AuthHandler struct {
	service   AuthService
	validator Validator
}

func NewAuthHandler(service AuthService, validator Validator) *AuthHandler {
	return &AuthHandler{
		service:   service,
		validator: validator,
	}
}

func (h *AuthHandler) SignUP(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "incorrect data format", http.StatusBadRequest)
		return
	}

	err := h.validator.ValidateStruct(&user)
	if err != nil {
		log.Error().Err(err).Msg("Validation failed")
		http.Error(w, "incorrect data for registration", http.StatusBadRequest)
		return
	}

	err = h.service.SignUP(&user)
	if err != nil {
		log.Error().Err(err).Msg("Error to signUP user")
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(map[string]string{
		"message": "User successfully registered"})
	if err != nil {
		http.Error(w, "could not encode response", http.StatusInternalServerError)
	}
}

func (h *AuthHandler) SignIN(w http.ResponseWriter, r *http.Request) {
	var user *model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Error().Err(err).Msg("Decoder in authHandler")
		http.Error(w, "incorrect data format", http.StatusBadRequest)
		return
	}

	token, err := h.service.SignIN(user.Email, user.Password)
	if err != nil {
		log.Error().Err(err).Msg("SignIN in authHandler")
		http.Error(w, "incorrect data format", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	err = json.NewEncoder(w).Encode(map[string]string{
		"token": token})
	if err != nil {
		return
	}
}
