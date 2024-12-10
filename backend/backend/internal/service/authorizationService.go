package service

import (
	"e-commerce/backend/internal/model"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
	"time"
)

type AuthRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	CreateWallet(wallet *model.Wallet) error
}

type UserRepositoryA interface {
	UserSave(user *model.User) error
}

type AuthService struct {
	repo  AuthRepository
	uRepo UserRepositoryA
}

func NewAuthService(repo AuthRepository, uRepo UserRepositoryA) *AuthService {
	return &AuthService{
		repo:  repo,
		uRepo: uRepo,
	}
}

func (s *AuthService) FindByEmail(email string) (*model.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *AuthService) GetAll(db *gorm.DB) (*model.User, error) {
	var users *model.User
	err := db.Model(&model.User{}).Preload("Wallet").Find(&users).Error
	return users, err
}

func (s *AuthService) SignUP(username, email, password, phone string) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("failed to hash password")
	}

	user := &model.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
		Phone:    phone,
	}
	if err := s.repo.Create(user); err != nil {
		log.Error().Err(err).Msg("Failed to signUP")
		return fmt.Errorf("failed to sign up: %v", err)
	}

	wallet := &model.Wallet{
		WalletID: user.UserID,
	}

	if err := s.repo.CreateWallet(wallet); err != nil {
		log.Error().Err(err).Msg("Failed to create wallet")
		return fmt.Errorf("failed to create wallet: %v", err)
	}

	user.WalletID = wallet.WalletID
	if err := s.uRepo.UserSave(user); err != nil {
		log.Error().Err(err).Msg("Error to update user with walletID")
		return fmt.Errorf("failed to update user with walletID: %v", err)
	}

	return nil
}

func (s *AuthService) SignIN(email, password string) (string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", fmt.Errorf("failed to find account %s: %v", email, err)
	}
	if user == nil {
		return "", fmt.Errorf("account %s is not found", email)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("incorrect password")
	}

	token, err := s.GenerateJWTToken(user.UserID, user.Role)
	if err != nil {
		log.Error().Err(err).Str("userID", fmt.Sprintf("%d", user.UserID)).Msg("Error generating token")
		return "", err
	}
	return token, nil
}

func (s *AuthService) GenerateJWTToken(userID int, role string) (string, error) {
	keyJWT := s.GetJWTKey()

	claims := jwt.MapClaims{
		"role": role,
		"sub":  userID,
		"exp":  time.Now().Add(time.Hour * 17).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(keyJWT))
	if err != nil {
		log.Error().Msg("error creating token")
		return "", fmt.Errorf("unable to create token: %v", err)
	}
	return tokenString, nil
}

func (s *AuthService) GetJWTKey() string {
	keyJWT := os.Getenv("JWT_SECRET_KEY")
	fmt.Println(len(keyJWT))
	if keyJWT == "" {
		log.Fatal().Msg("JWT secret key is not set")
	}
	return keyJWT
}
