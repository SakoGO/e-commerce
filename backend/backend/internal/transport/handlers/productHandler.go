package handlers

import (
	"e-commerce/backend/internal/model"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type ProductService interface {
	CreateProduct(ownerID, ShopID int, name, description, price, stock, image string) error
	GetProductByID(productID int) (*model.Product, error)
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

	err = h.ProductService.CreateProduct(ownerID, product.ShopID, product.Name, product.Description, product.Price, product.Stock, product.Image)
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
		return
	}

}

func getProductID(w http.ResponseWriter, r *http.Request) (int, error) {
	idStr := chi.URLParam(r, "id")
	productID, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error().Err(err).Msg("Invalid id format")
		http.Error(w, "invalid id format", http.StatusBadRequest)
		return 0, err
	}
	return productID, nil
}

func (h *Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {

	productID, err := getProductID(w, r)
	if err != nil {
		return
	}

	product, err := h.ProductService.GetProductByID(productID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error().Err(err).Msg("product by id not found")
			http.Error(w, "product not found", http.StatusNotFound)
			return
		}
		http.Error(w, "error get product", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&product); err != nil {
		log.Error().Err(err).Msg("Error encoding data")
		http.Error(w, "error encoding data", http.StatusInternalServerError)
		return
	}
	log.Info().Msgf("product %d sucesfully encoded")
}
