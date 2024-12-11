package repository

import (
	"e-commerce/backend/internal/model"
	"gorm.io/gorm"
)

type ShopRepository struct {
	db *gorm.DB
}

func NewShopRepository(db *gorm.DB) (*ShopRepository, error) {
	err := db.AutoMigrate(&model.Shop{})
	if err != nil {
		return nil, err
	}
	return &ShopRepository{db: db}, nil
}

func (r *ShopRepository) CreateShop(shop *model.Shop) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		err := r.db.Create(shop).Error
		if err != nil {
			return err
		}
		err = r.db.Model(&model.User{}).Where("user_id = ?", shop.OwnerID).Update("has_shop", true).Error
		if err != nil {
			return err
		}
		return nil
	})
}
