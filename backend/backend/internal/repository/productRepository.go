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

func (r *ProductRepository) GetProductByID(productID int) (*model.Product, error) {
	var product model.Product
	err := r.db.First(&product, productID).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) GetProductsByShopID(shopID int) ([]model.Product, error) {
	var products []model.Product
	err := r.db.Where("shop_id = ?", shopID).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) UpdateProduct(product *model.Product) error {
	var update model.Product
	err := r.db.First(&update, product.ProductID).Error
	if err != nil {
		return err
	}

	update.Name = product.Name
	update.Description = product.Description
	update.Price = product.Price
	update.Stock = product.Stock
	update.Image = product.Image
	update.CategoryID = product.CategoryID

	return r.db.Save(&update).Error
}

func (r *ProductRepository) DeleteProduct(productID int) error {
	var product model.Product

	err := r.db.Where("product_id = ?", productID).Delete(&product).Error
	if err != nil {
		return err
	}
	return nil

}
