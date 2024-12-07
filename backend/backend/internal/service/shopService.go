package service

import (
	"e-commerce/backend/internal/model"
)

type ShopRepository interface {
	Create(shop *model.Shop, ownerID int) error
	GetByID(id, ownerID int) (*model.Shop, error)
	GetAll(ownerID int) ([]model.Shop, error)
	Update(shop *model.Shop, ownerID int) error
	Delete(id, ownerID int) error
}

type ShopService struct {
	repo ShopRepository
}

func NewShopService(repo ShopRepository) *ShopService {
	return &ShopService{repo: repo}
}

func (s *ShopService) Create(shop *model.Shop, ownerID int) error {
	return s.repo.Create(shop, ownerID)
}
