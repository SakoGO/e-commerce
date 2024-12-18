package handlers

import (
	"e-commerce/backend/internal/model"
	"encoding/json"
	"errors"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"net/http"
)

type ShopService interface {
	CreateShop(userID int, shop *model.Shop) error
	GetShopID(shopID int) (*model.Shop, error)
	UpdateShop(ownerID int, shop *model.Shop) error
	DeleteShop(shopID, ownerID int) error
}

type ShopHandler struct {
	service   ShopService
	validator Validator
}

func NewShopHandler(service ShopService, validator Validator) *ShopHandler {
	return &ShopHandler{
		service:   service,
		validator: validator,
	}
}

func (h *ShopHandler) CreateShop(w http.ResponseWriter, r *http.Request) {
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

	err = h.service.CreateShop(userID, &shop)
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

func (h *ShopHandler) GetShopID(w http.ResponseWriter, r *http.Request) {

	shopID, err := GetID(w, r)
	if err != nil {
		return
	}

	shop, err := h.service.GetShopID(shopID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error().Err(err).Msg("shop by id not found")
			http.Error(w, "shop not found", http.StatusNotFound)
			return
		}
		http.Error(w, "error get shop", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&shop); err != nil {
		log.Error().Err(err).Msg("error encoding data")
		http.Error(w, "error encoding data", http.StatusInternalServerError)
		return
	}
	log.Info().Msgf("shop %d sucesfully encoded")
}

func (h *ShopHandler) UpdateShop(w http.ResponseWriter, r *http.Request) {
	var shop model.Shop
	if err := json.NewDecoder(r.Body).Decode(&shop); err != nil {
		http.Error(w, "error decodeing data", http.StatusBadRequest)
		return
	}

	ownerID, ok := r.Context().Value("userID").(int)
	if !ok {
		log.Error().Msg("User ID not found in context")
		http.Error(w, "user not authorized", http.StatusUnauthorized)
		return
	}

	err := h.service.UpdateShop(ownerID, &shop)
	if err != nil {
		log.Error().Err(err).Msg("Error updating shop")
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
}

func (h *ShopHandler) DeleteShop(w http.ResponseWriter, r *http.Request) {
	shopID, err := GetID(w, r)
	if err != nil {
		return
	}

	ownerID, ok := r.Context().Value("userID").(int)
	if !ok {
		log.Error().Msg("User ID not found in context")
		http.Error(w, "user not authorized", http.StatusUnauthorized)
		return
	}

	err = h.service.DeleteShop(shopID, ownerID)
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
		return
	}
}
