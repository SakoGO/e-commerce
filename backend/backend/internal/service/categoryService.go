package service

import (
	"e-commerce/backend/internal/model"
	"fmt"
	"github.com/rs/zerolog/log"
)

type CategoryRepository interface {
	CreateCategory(category *model.Category) error
	GetCategoryByID(categoryID int) (*model.Category, error)
	GetProductsByCategoryID(categoryID int) ([]model.Product, error)
}

type CategoryService struct {
	cRepo CategoryRepository
}

func NewCategoryService(cRepo CategoryRepository) *CategoryService {
	return &CategoryService{
		cRepo: cRepo,
	}
}

func (s *CategoryService) CreateCategory(name string) error {

	category := &model.Category{Name: name}

	if err := s.cRepo.CreateCategory(category); err != nil {
		log.Error().Err(err).Msg("Failed to create category")
		return fmt.Errorf("failed to create category: %v", err)
	}
	return nil
}

func (c CategoryService) GetCategoryByID(categoryID int) (*model.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryService) GetProductsByCategoryID(categoryID int) ([]model.Product, error) {
	//TODO implement me
	panic("implement me")
}
