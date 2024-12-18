package service

import (
	"e-commerce/backend/internal/model"
	"fmt"
	"github.com/rs/zerolog/log"
)

type ShopRepository interface {
	CreateShop(shop *model.Shop) error
	UpdateShop(shop *model.Shop) error
	DeleteShop(shopID, ownerID int) error
	GetShopID(shopID int) (*model.Shop, error)
}

type ShopService struct {
	sRepo ShopRepository
	uRepo UserRepository
}

func NewShopService(sRepo ShopRepository, uRepo UserRepository) *ShopService {
	return &ShopService{
		sRepo: sRepo,
		uRepo: uRepo,
	}
}

func (s *ShopService) CreateShop(userID int, shop *model.Shop) error {

	user, err := s.uRepo.UserFindByID(userID)
	if err != nil {
		log.Error().Msg("Failed to create shop")
		return fmt.Errorf("failed to create shop: %v", err)
	}

	shop.OwnerID = userID

	if user.HasShop {
		log.Error().Msg("User already has shop")
		return fmt.Errorf("You already has shop")
	}

	if err := s.sRepo.CreateShop(shop); err != nil {
		log.Error().Err(err).Msg("Failed to create shop")
		return fmt.Errorf("failed to create shop: %v", err)
	}

	return nil
}

func (s *ShopService) GetShopID(shopID int) (*model.Shop, error) {
	shop, err := s.sRepo.GetShopID(shopID)
	if err != nil {
		return nil, err
	}
	return shop, nil
}

func (s *ShopService) UpdateShop(ownerID int, shop *model.Shop) error {

	Shop, err := s.sRepo.GetShopID(shop.ShopID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to find shop by id")
		return fmt.Errorf("failed to find shop by id: %v", err)
	}

	if Shop.OwnerID != ownerID {
		log.Error().Msg("The owner ID in the request does not match the actual one")
		return fmt.Errorf("the owner ID in the request does not match the actual one")
	}

	if err := s.sRepo.UpdateShop(shop); err != nil {
		log.Error().Err(err).Msg("Failed to update shop")
		return fmt.Errorf("failed to update shop: %v", err)
	}
	return nil
}

func (s *ShopService) DeleteShop(shopID, ownerID int) error {

	Shop, err := s.sRepo.GetShopID(shopID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to find shop by id")
		return fmt.Errorf("failed to find shop by id: %v", err)
	}

	if Shop.OwnerID != ownerID {
		log.Error().Msg("The owner ID in the request does not match the actual one")
		return fmt.Errorf("the owner ID in the request does not match the actual one")
	}

	if err := s.sRepo.DeleteShop(shopID, ownerID); err != nil {
		log.Error().Err(err).Msg("Failed to delete shop")
		return fmt.Errorf("failed to delete shop: %v", err)
	}

	return nil
}
