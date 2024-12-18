package service

import (
	"e-commerce/backend/internal/model"
	"fmt"
	"github.com/rs/zerolog/log"
)

type ProductRepository interface {
	CreateProduct(product *model.Product) error
	GetProductByID(productID int) (*model.Product, error)
	GetProductsByShopID(shopID int) ([]model.Product, error)
	UpdateProduct(product *model.Product) error
	DeleteProduct(productID int) error
	DeleteProductsByShopID(shopID int) error
}

type ProductService struct {
	pRepo ProductRepository
	sRepo ShopRepository
}

func NewProductService(pRepo ProductRepository, sRepo ShopRepository) *ProductService {
	return &ProductService{
		pRepo: pRepo,
		sRepo: sRepo,
	}
}

func (s *ProductService) CreateProduct(ownerID int, product *model.Product) error {

	shop, err := s.sRepo.GetShopID(product.ShopID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to find shop by id")
		return fmt.Errorf("failed to find shop by id: %v", err)
	}

	if shop.OwnerID != ownerID {
		log.Error().Msg("The owner ID in the request does not match the actual one")
		return fmt.Errorf("the owner ID in the request does not match the actual one")
	}

	if err := s.pRepo.CreateProduct(product); err != nil {
		log.Error().Err(err).Msg("Failed to create product")
		return fmt.Errorf("failed to create product: %v", err)
	}
	return nil
}

func (s *ProductService) GetProductByID(productID int) (*model.Product, error) {
	product, err := s.pRepo.GetProductByID(productID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) GetProductsByShopID(shopID int) ([]model.Product, error) {
	products, err := s.pRepo.GetProductsByShopID(shopID)
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		log.Error().Err(err).Msg("no products found for request shop")
		return nil, fmt.Errorf("no products found for shop with ID %d", shopID)
	}

	return products, nil
}

func (s *ProductService) UpdateProduct(ownerID int, product *model.Product) error {

	shop, err := s.sRepo.GetShopID(product.ShopID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to find shop by id")
		return fmt.Errorf("failed to find shop by id: %v", err)
	}

	if shop.OwnerID != ownerID {
		log.Error().Msg("The owner ID in the request does not match the actual one")
		return fmt.Errorf("the owner ID in the request does not match the actual one")
	}

	if err := s.pRepo.UpdateProduct(product); err != nil {
		log.Error().Err(err).Msg("Failed to update shop")
		return fmt.Errorf("failed to update shop: %v", err)
	}
	return nil
}

func (s *ProductService) DeleteProductsByShopID(shopID, ownerID int) error {
	shop, err := s.sRepo.GetShopID(shopID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to find shop by id")
		return fmt.Errorf("failed to find shop by id: %v", err)
	}

	if shop.OwnerID != ownerID {
		log.Error().Msg("The owner ID in the request does not match the actual one")
		return fmt.Errorf("the owner ID in the request does not match the actual one")
	}

	err = s.pRepo.DeleteProductsByShopID(shopID)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) DeleteProduct(productID, ownerID, shopID int) error {

	shop, err := s.sRepo.GetShopID(shopID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to find shop by id")
		return fmt.Errorf("failed to find shop by id: %v", err)
	}

	if shop.OwnerID != ownerID {
		log.Error().Msg("The owner ID in the request does not match the actual one")
		return fmt.Errorf("the owner ID in the request does not match the actual one")
	}

	if err := s.pRepo.DeleteProduct(productID); err != nil {
		log.Error().Err(err).Msg("Failed to delete product")
		return fmt.Errorf("failed to delete shop: %v", err)
	}

	return nil
}
