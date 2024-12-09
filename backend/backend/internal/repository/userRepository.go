package repository

import (
	model "e-commerce/backend/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) (*UserRepository, error) {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}
	return &UserRepository{db: db}, nil
}

func (r *UserRepository) UserFindByID(userID int) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UserUpdate(user *model.User) error {
	/*var update model.User
	err := r.db.First(&update, user.UserID).Error
	if err != nil {
		return err
	}

	update.Username = user.Username
	update.Email = user.Email
	update.Password = user.Password
	update.Phone = user.Phone
	update.WalletID = user.WalletID */

	return r.db.Save(user).Error
}

/*
func (r *UserRepository) UserDelete(userID int) error {
	var user model.User
	err := r.db.First(&user, userID).Error
	if err != nil {
		return err
	}

	err = r.db.Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}*/
