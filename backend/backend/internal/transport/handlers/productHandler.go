package handlers

import (
	"e-commerce/backend/internal/model"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

type ProductService interface {
	CreateProduct(ownerID int, name, description, price, stock, image string) error
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "incorrect data format", http.StatusBadRequest)
		return
	}

	err := h.validator.ValidateStruct(&product)
	if err != nil {
		log.Error().Err(err).Msg("Validation failed")
		http.Error(w, "incorrect data for creating shop", http.StatusBadRequest)
		return
	}

	ownerID, ok := r.Context().Value("userID").(int)
	if !ok {
		log.Error().Msg("User ID not found in context")
		http.Error(w, "user not authorized", http.StatusUnauthorized)
		return
	}

	err = h.ProductService.CreateProduct(ownerID, product.Name, product.Description, product.Price, product.Stock, product.Image)
	if err != nil {
		log.Error().Err(err).Msg("Error creating product")
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(map[string]string{
		"message": "Product succesfully created"})
	if err != nil {
		http.Error(w, "could not encode response", http.StatusInternalServerError)
	}

}
