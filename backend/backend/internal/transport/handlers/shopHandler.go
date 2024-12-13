package handlers

import (
	"e-commerce/backend/internal/model"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

type ShopService interface {
	CreateShop(userID int, name, email, description string) error
	UpdateShop(shopID, ownerID int, name, description, email string) error
	DeleteShop(shopID, ownerID int) error
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

func (h *Handler) UpdateShop(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	shopID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id format", http.StatusBadRequest)
		return
	}

	ownerID, ok := r.Context().Value("userID").(int)
	if !ok {
		log.Error().Msg("User ID not found in context")
		http.Error(w, "user not authorized", http.StatusUnauthorized)
		return
	}

	var shop model.Shop
	if err := json.NewDecoder(r.Body).Decode(&shop); err != nil {
		http.Error(w, "error decodeing data", http.StatusBadRequest)
		return
	}

	err = h.ShopService.UpdateShop(shopID, ownerID, shop.Name, shop.Description, shop.Email)
	if err != nil {
		log.Error().Err(err).Msg("Error updating shop")
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
}

func (h *Handler) DeleteShop(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	shopID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id format", http.StatusBadRequest)
		return
	}

	ownerID, ok := r.Context().Value("userID").(int)
	if !ok {
		log.Error().Msg("User ID not found in context")
		http.Error(w, "user not authorized", http.StatusUnauthorized)
		return
	}

	err = h.ShopService.DeleteShop(shopID, ownerID)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting shop")
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{
		"message": "Shop succesfully deleted"})
	if err != nil {
		http.Error(w, "could not encode response", http.StatusInternalServerError)
	}

}
