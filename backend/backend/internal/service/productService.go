package service

import (
	"e-commerce/backend/internal/model"
	"fmt"
	"github.com/rs/zerolog/log"
)

type ProductRepository interface {
	CreateProduct(product *model.Product) error
	GetShopByID(id int) (*model.Shop, error)
}

type ProductService struct {
	repo ProductRepository
}

func NewProductRepository(repo ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) GetShopByID(shopID int) (*model.Shop, error) {
	return s.repo.GetShopByID(shopID)
}

func (s *ProductService) CreateProduct(ownerID, shopID int, name, description, price, stock, image string) error {

	shop, err := s.repo.GetShopByID(shopID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to find shop by id")
		return fmt.Errorf("failed to find shop by id: %v", err)
	}

	if shop.OwnerID != ownerID {
		log.Error().Msg("The owner ID in the request does not match the actual one")
		return fmt.Errorf("the owner ID in the request does not match the actual one")
	}

	product := &model.Product{
		ShopID:      shopID,
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
