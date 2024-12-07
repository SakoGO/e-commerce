package repository

import (
	"e-commerce/backend/internal/model"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func (r *authRepository) SignIN(user *model.User) error {
	return r.db.Where(user).Error
}

func NewAuthRepository(db *gorm.DB) (*authRepository, error) {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}
	return &authRepository{db: db}, nil
}

func (r *authRepository) SignUP(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *authRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) FindByPhone(phone string) (*model.User, error) {
	var user model.User
	err := r.db.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
