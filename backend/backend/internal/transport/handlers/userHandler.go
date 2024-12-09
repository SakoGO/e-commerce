package handlers

import (
	"e-commerce/backend/internal/model"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

type UserService interface {
	UserUpdate(userID int, user *model.User) (*model.User, error)
}

func (h *Handler) UserUpdate(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil || userID <= 0 {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	tokenUserID := r.Context().Value("userID").(int)

	if userID != tokenUserID {
		http.Error(w, "Error updating data, incorrect ID", http.StatusForbidden)
		return
	}

	var userData model.User
	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		http.Error(w, fmt.Sprintf("Invalid input data: %v", err), http.StatusBadRequest)
		return
	}

	updatedUser, err := h.UserService.UserUpdate(userID, &userData)
	if err != nil {
		log.Error().Err(err).Msg("failed to update user")
		http.Error(w, "failed to update user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(updatedUser); err != nil {
		log.Error().Err(err).Msg("failed to encode response")
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
