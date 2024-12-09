package service

import "e-commerce/backend/internal/model"

type WalletRepository interface {
	WalletBalance(balance float64) (*model.Wallet, error)
	CreateWallet(wallet *model.Wallet) error
	FindByUserID(userID int) (*model.Wallet, error)
}

type WalletService struct {
	repo WalletRepository
}

func NewWalletService(repo WalletRepository) *WalletService {
	return &WalletService{repo: repo}
}

func (s *WalletService) WalletBalance(balance float64) (*model.Wallet, error) {
	return s.repo.WalletBalance(balance)
}

func (s *WalletService) CreateWallet(wallet *model.Wallet) error {
	return s.repo.CreateWallet(wallet)
}

func (s *WalletService) FindByUserID(userID int) (*model.Wallet, error) {
	return s.repo.FindByUserID(userID)
}
