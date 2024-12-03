package handlers

import (
	"e-commerce/backend/internal/model"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"net/http"
)

type UserService interface {
	SignUP(username, email, password, phone string) error
	SignIN(email, password string) (string, error)
	UserFindByUsername(username string) (*model.User, error)
	UserFindByEmail(email string) (*model.User, error)
	UserFindByID(userID int) (*model.User, error) //
	UserDelete(userID int) error                  //
}

type Validator interface {
	ValidateStruct(interface{}) error
}

func (h *Handler) SignUP(w http.ResponseWriter, r *http.Request) {
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

	err = h.UserService.SignUP(user.Username, user.Email, user.Password, user.Phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(map[string]string{
		"-": "User successfully registered"})
	if err != nil {
		http.Error(w, "could not encode response", http.StatusInternalServerError)
	}
}

func (h *Handler) SignIN(w http.ResponseWriter, r *http.Request) {
	var user *model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Error().Err(err).Msg("Гандон 1")
		http.Error(w, "incorrect data format", http.StatusBadRequest)
		return
	}

	token, err := h.UserService.SignIN(user.Email, user.Password)
	if err != nil {
		log.Error().Err(err).Msg("Гандон 2")
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

// ///////////////////////////////////////////////////////////////////////////////////////////////////
func (h *Handler) UserFindByUsername(w http.ResponseWriter, r *http.Request) {
	// Извлекаем username из URL-параметра
	username := chi.URLParam(r, "username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	// Ищем пользователя по имени
	user, err := h.UserService.UserFindByUsername(username)
	if err != nil {
		// Если ошибка поиска (например, пользователь не найден)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Кодируем пользователя в JSON и отправляем в ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
