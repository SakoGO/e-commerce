package repository

import (
	"e-commerce/backend/internal/model"
	"gorm.io/gorm"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) (*WalletRepository, error) {
	err := db.AutoMigrate(&model.Wallet{})
	if err != nil {
		return nil, err
	}
	return &WalletRepository{db: db}, nil
}

func (r *WalletRepository) CreateWallet(wallet *model.Wallet) error {
	return r.db.Create(wallet).Error
}

func (r *WalletRepository) FindByUserID(userID int) (*model.Wallet, error) {
	var wallet model.Wallet
	err := r.db.Where("user_id = ?", userID).First(&wallet).Error
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *WalletRepository) WalletBalance(balance float64) (*model.Wallet, error) {
	var wallet model.Wallet
	err := r.db.First(&wallet, balance).Error
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}
