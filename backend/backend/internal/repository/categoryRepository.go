package repository

import (
	"e-commerce/backend/internal/model"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) (*CategoryRepository, error) {
	err := db.AutoMigrate(&model.Category{})
	if err != nil {
		return nil, err
	}
	return &CategoryRepository{db: db}, nil
}

func (r *CategoryRepository) CreateCategory(category *model.Category) error {
	return r.db.Create(category).Error
}

func (r *CategoryRepository) GetCategoryByID(categoryID int) (*model.Category, error) {
	var category model.Category
	err := r.db.First(&category, categoryID).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) GetProductsByCategoryID(categoryID int) ([]model.Product, error) {
	var products []model.Product
	err := r.db.Where("category_id = ?", categoryID).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
