package service

import (
	"e-commerce/backend/internal/model"
	"fmt"
	"github.com/rs/zerolog/log"
)

type ProductRepository interface {
	CreateProduct(product *model.Product) error
}

type ProductService struct {
	repo ProductRepository
}

func NewProductRepository(repo ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(ownerID int, name, description, price, stock, image string) error {

	product := &model.Product{
		SellerID:    ownerID,
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
		Image:       image,
		Category:    model.Category{},
	}

	if err := s.repo.CreateProduct(product); err != nil {
		log.Error().Err(err).Msg("Failed to create product")
		return fmt.Errorf("failed to create product: %v", err)
	}
	return nil
}
