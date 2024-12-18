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

func (r *ShopRepository) GetShopID(shopID int) (*model.Shop, error) {
	var shop model.Shop
	err := r.db.First(&shop, shopID).Error
	if err != nil {
		return nil, err
	}
	return &shop, nil
}

func (r *ShopRepository) UpdateShop(shop *model.Shop) error {
	var update model.Shop
	err := r.db.First(&update, shop.ShopID).Error
	if err != nil {
		return err
	}

	update.Name = shop.Name
	update.Description = shop.Description
	update.Email = shop.Email

	return r.db.Save(&update).Error
}

func (r *ShopRepository) DeleteShop(shopID, ownerID int) error {
	return r.db.Transaction(func(tx *gorm.DB) error {

		//1 Удаляем все продукты по shopID
		//2 Удаляем шоп
		//3 Ставим пользователю hasShop = 0
		err := r.db.Where("shop_id = ?", shopID).Delete(&model.Product{}).Error
		if err != nil {
			return err
		}

		err = r.db.Where("shop_id = ?", shopID).Delete(&model.Shop{}).Error
		if err != nil {
			return err
		}

		err = r.db.Model(&model.User{}).Where("user_id = ?", ownerID).Update("has_shop", false).Error
		if err != nil {
			return err
		}

		return nil
	})
}
