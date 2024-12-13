package repository

import (
	"e-commerce/backend/internal/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) (*ProductRepository, error) {
	err := db.AutoMigrate(&model.Product{})
	if err != nil {
		return nil, err
	}
	return &ProductRepository{db: db}, nil
}

func (r *ProductRepository) CreateProduct(product *model.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) GetShopByID(shopID int) (*model.Shop, error) {
	var shop model.Shop
	err := r.db.First(&shop, shopID).Error
	if err != nil {
		return nil, err
	}
	return &shop, nil
}
