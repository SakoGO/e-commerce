package handlers

import (
	"e-commerce/backend/internal/model"
	"encoding/json"
	"errors"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"net/http"
)

type ProductService interface {
	CreateProduct(ownerID int, product *model.Product) error
	GetProductByID(productID int) (*model.Product, error)
	GetProductsByShopID(shopID int) ([]model.Product, error)
	UpdateProduct(ownerID int, product *model.Product) error
	DeleteProduct(productID, ownerID, shopID int) error
	DeleteProductsByShopID(shopID, ownerID int) error
}

type ProductHandler struct {
	service   ProductService
	validator Validator
}

func NewProductHandler(service ProductService, validator Validator) *ProductHandler {
	return &ProductHandler{
		service:   service,
		validator: validator,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
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

	err = h.service.CreateProduct(ownerID, &product)
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

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {

	productID, err := GetID(w, r)
	if err != nil {
		return
	}

	product, err := h.service.GetProductByID(productID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error().Err(err).Msg("product by id not found")
			http.Error(w, "product not found", http.StatusNotFound)
			return
		}
		http.Error(w, "error get product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&product); err != nil {
		log.Error().Err(err).Msg("Error encoding data")
		http.Error(w, "error encoding data", http.StatusInternalServerError)
		return
	}
}

func (h *ProductHandler) GetProductsByShopID(w http.ResponseWriter, r *http.Request) {
	shopID, err := GetID(w, r)
	if err != nil {
		return
	}

	products, err := h.service.GetProductsByShopID(shopID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error().Err(err).Msg("shop by id not found")
			http.Error(w, "shop not found", http.StatusNotFound)
			return
		}
		http.Error(w, "error get products", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&products); err != nil {
		log.Error().Err(err).Msg("Error encoding data")
		http.Error(w, "error encoding data", http.StatusInternalServerError)
		return
	}
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "incorrect data format", http.StatusBadRequest)
		return
	}

	//TODO:: Maybe add validation

	ownerID, ok := r.Context().Value("userID").(int)
	if !ok {
		log.Error().Msg("User ID not found in context")
		http.Error(w, "user not authorized", http.StatusUnauthorized)
		return
	}

	err := h.service.UpdateProduct(ownerID, &product)
	if err != nil {
		log.Error().Err(err).Msg("Error updating product")
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{
		"message": "Product succesfully updated"})
	if err != nil {
		http.Error(w, "could not encode response", http.StatusInternalServerError)
		return
	}
}

func (h *ProductHandler) DeleteProductsByShopID(w http.ResponseWriter, r *http.Request) {
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

	err = h.service.DeleteProductsByShopID(shopID, ownerID)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting products")
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{
		"message": "All products succesfully deleted"})
	if err != nil {
		http.Error(w, "could not encode response", http.StatusInternalServerError)
		return
	}
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := GetID(w, r)
	if err != nil {
		return
	}

	var product model.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "incorrect data format", http.StatusBadRequest)
		return
	}

	ownerID, ok := r.Context().Value("userID").(int)
	if !ok {
		log.Error().Msg("User ID not found in context")
		http.Error(w, "user not authorized", http.StatusUnauthorized)
		return
	}

	err = h.service.DeleteProduct(productID, ownerID, product.ShopID)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting product")
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{
		"message": "Product succesfully deleted"})
	if err != nil {
		http.Error(w, "could not encode response", http.StatusInternalServerError)
		return
	}
}
