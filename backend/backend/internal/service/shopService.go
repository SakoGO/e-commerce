package service

import (
	"e-commerce/backend/internal/model"
	"fmt"
	"github.com/rs/zerolog/log"
)

type ShopRepository interface {
	CreateShop(shop *model.Shop) error
	UpdateShop(shop *model.Shop) error
	DeleteShop(shopID int) error
	GetShopID(shopID int) (*model.Shop, error)
}

type ShopService struct {
	repo  ShopRepository
	uRepo UserRepository
}

func NewShopService(repo ShopRepository, uRepo UserRepository) *ShopService {
	return &ShopService{
		repo:  repo,
		uRepo: uRepo,
	}
}

func (s *ShopService) CreateShop(userID int, name, email, description string) error {

	shop := &model.Shop{
		Name:        name,
		Email:       email,
		Description: description,
		OwnerID:     userID,
		Products:    []*model.Product{},
	}

	user, err := s.uRepo.UserFindByID(userID)
	if err != nil {
		log.Error().Msg("Failed to create shop")
		return fmt.Errorf("failed to create shop: %v", err)
	}

	if user.HasShop {
		log.Error().Msg("User already has shop")
		return fmt.Errorf("You already has shop")
	}

	if err := s.repo.CreateShop(shop); err != nil {
		log.Error().Err(err).Msg("Failed to create shop")
		return fmt.Errorf("failed to create shop: %v", err)
	}

	return nil
}

func (s *ShopService) UpdateShop(shopID, ownerID int, name, description, email string) error {

	Shop, err := s.repo.GetShopID(shopID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to find shop by id")
		return fmt.Errorf("failed to find shop by id: %v", err)
	}

	if Shop.OwnerID != ownerID {
		log.Error().Msg("The owner ID in the request does not match the actual one")
		return fmt.Errorf("the owner ID in the request does not match the actual one")
	}

	updatedShop := &model.Shop{
		ShopID:      shopID,
		Name:        name,
		Description: description,
		Email:       email,
		OwnerID:     ownerID,
	}

	if err := s.repo.UpdateShop(updatedShop); err != nil {
		log.Error().Err(err).Msg("Failed to update shop")
		return fmt.Errorf("failed to update shop: %v", err)
	}
	return nil
}

func (s *ShopService) DeleteShop(shopID, ownerID int) error {

	Shop, err := s.repo.GetShopID(shopID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to find shop by id")
		return fmt.Errorf("failed to find shop by id: %v", err)
	}

	if Shop.OwnerID != ownerID {
		log.Error().Msg("The owner ID in the request does not match the actual one")
		return fmt.Errorf("the owner ID in the request does not match the actual one")
	}

	if err := s.repo.DeleteShop(shopID); err != nil {
		log.Error().Err(err).Msg("Failed to delete shop")
		return fmt.Errorf("failed to delete shop: %v", err)
	}

	return nil
}
