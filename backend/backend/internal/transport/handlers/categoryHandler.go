package handlers

import (
	"e-commerce/backend/internal/model"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

type CategoryService interface {
	CreateCategory(name string) error
}

func (h *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category model.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, "incorrect data format", http.StatusBadRequest)
		return
	}

	err := h.CategoryService.CreateCategory(category.Name)
	if err != nil {
		log.Error().Err(err).Msg("Error creating category")
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(map[string]string{
		"message": "Category succesfully created"})
	if err != nil {
		http.Error(w, "could not encode response", http.StatusInternalServerError)
		return
	}
}
