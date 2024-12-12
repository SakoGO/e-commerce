package service

import (
	"e-commerce/backend/internal/model"
	"fmt"
	"github.com/rs/zerolog/log"
)

type ShopRepository interface {
	CreateShop(shop *model.Shop) error
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
