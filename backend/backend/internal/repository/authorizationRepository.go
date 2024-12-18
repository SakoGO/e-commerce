package repository

import (
	"e-commerce/backend/internal/model"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) (*AuthRepository, error) {
	err := db.AutoMigrate(&model.User{}, &model.Wallet{})
	if err != nil {
		return nil, err
	}
	return &AuthRepository{db: db}, nil
}

func (r *AuthRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

/*
func (r *AuthRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
*/

func (r *AuthRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

/*
func (r *AuthRepository) FindByPhone(phone string) (*model.User, error) {
	var user model.User
	err := r.db.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
*/
