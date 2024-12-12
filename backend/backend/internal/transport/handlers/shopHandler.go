package handlers

import (
	"e-commerce/backend/internal/model"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

type ShopService interface {
	CreateShop(userID int, name, email, description string) error
}

func (h *Handler) CreateShop(w http.ResponseWriter, r *http.Request) {
	var shop model.Shop
	if err := json.NewDecoder(r.Body).Decode(&shop); err != nil {
		http.Error(w, "incorrect data format", http.StatusBadRequest)
		return
	}

	err := h.validator.ValidateStruct(&shop)
	if err != nil {
		log.Error().Err(err).Msg("Validation failed")
		http.Error(w, "incorrect data for creating shop", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		log.Error().Msg("User ID not found in context")
		http.Error(w, "user not authorized", http.StatusUnauthorized)
		return
	}

	err = h.ShopService.CreateShop(userID, shop.Name, shop.Email, shop.Description)
	if err != nil {
		log.Error().Err(err).Msg("Error creating shop")
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(map[string]string{
		"message": "Shop succesfully created"})
	if err != nil {
		http.Error(w, "could not encode response", http.StatusInternalServerError)
	}
}
