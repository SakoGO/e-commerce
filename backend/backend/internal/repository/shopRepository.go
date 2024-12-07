package repository

import (
	"e-commerce/backend/internal/model"
	"errors"
	"gorm.io/gorm"
)

type ShopRepository struct {
	db *gorm.DB
}

func NewShopRepository(db *gorm.DB) *ShopRepository {
	return &ShopRepository{db: db}
}

func (r *ShopRepository) Create(shop *model.Shop, ownerID int) error {
	shop.OwnerID = ownerID
	return r.db.Create(shop).Error
}

func (r *ShopRepository) GetByID(id, ownerID int) (*model.Shop, error) {
	var shop model.Shop
	if err := r.db.Preload("Products").Preload("Owner").Where("id = ? AND owner_id = ?", id, ownerID).First(&shop).Error; err != nil {
		return nil, err
	}
	return &shop, nil
}

func (r *ShopRepository) GetAll(ownerID int) ([]model.Shop, error) {
	var shops []model.Shop
	if err := r.db.Preload("Owner").Where("owner_id = ?", ownerID).Find(&shops).Error; err != nil {
		return nil, err
	}
	return shops, nil
}

func (r *ShopRepository) Update(shop *model.Shop, ownerID int) error {
	if err := r.db.Where("id = ? AND owner_id = ?", shop.ShopID, ownerID).First(shop).Error; err != nil {
		return errors.New("shop not found or not owned by this user")
	}
	return r.db.Save(shop).Error
}

func (r *ShopRepository) Delete(id, ownerID int) error {
	var shop model.Shop
	if err := r.db.Where("id = ? AND owner_id = ?", id, ownerID).First(&shop).Error; err != nil {
		return errors.New("shop not found or not owned by this user")
	}
	return r.db.Delete(&shop).Error
}
