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

func (s *ProductService) CreateProduct(ownerID, shopID, categoryID int, name, description, price, stock, image string) error {

	shop, err := s.sRepo.GetShopID(shopID)
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
		CategoryID:  categoryID,
		Category:    model.Category{},
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

func (s *ProductService) UpdateProduct(productID, ownerID, shopID, categoryID int, name, description, price, stock, image string) error {
	/*	product, err := s.pRepo.GetProductByID(productID)
		if err != nil {
			log.Error().Err(err).Msg("Failed to find product by ID")
			return fmt.Errorf("failed to find product by id: %v", err)
		}
	*/
	shop, err := s.sRepo.GetShopID(shopID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to find shop by id")
		return fmt.Errorf("failed to find shop by id: %v", err)
	}
	/*
		if product.ShopID != shopID {
			log.Error().Msg("The shop ID in the request does not match the actual one")
			return fmt.Errorf("the shop ID in the request does not match the actual one")
		}
	*/
	if shop.OwnerID != ownerID {
		log.Error().Msg("The owner ID in the request does not match the actual one")
		return fmt.Errorf("the owner ID in the request does not match the actual one")
	}

	updatedProduct := &model.Product{
		ProductID:   productID,
		ShopID:      shopID,
		CategoryID:  categoryID,
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
		Image:       image,
	}

	if err := s.pRepo.UpdateProduct(updatedProduct); err != nil {
		log.Error().Err(err).Msg("Failed to update shop")
		return fmt.Errorf("failed to update shop: %v", err)
	}
	return nil
}

func (s *ProductService) DeleteProduct(productID, ownerID, shopID int) error {
	product, err := s.pRepo.GetProductByID(productID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to find product by ID")
		return fmt.Errorf("failed to find product by id: %v", err)
	}

	shop, err := s.sRepo.GetShopID(shopID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to find shop by id")
		return fmt.Errorf("failed to find shop by id: %v", err)
	}

	if product.ShopID != shopID {
		log.Error().Msg("The shop ID in the request does not match the actual one")
		return fmt.Errorf("the shop ID in the request does not match the actual one")
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
