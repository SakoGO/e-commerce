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
}

type AuthService struct {
	aRepo AuthRepository
}

func NewAuthService(aRepo AuthRepository) *AuthService {
	return &AuthService{
		aRepo: aRepo,
	}
}

func (s *AuthService) FindByEmail(email string) (*model.User, error) {
	return s.aRepo.FindByEmail(email)
}

func (s *AuthService) GetAll(db *gorm.DB) (*model.User, error) {
	var users *model.User
	err := db.Model(&model.User{}).Preload("Wallet").Find(&users).Error
	return users, err
}

func (s *AuthService) SignUP(user *model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("failed to hash password")
	}

	user.Password = string(hashedPassword)
	user.Wallet = &model.Wallet{}

	if err := s.aRepo.Create(user); err != nil {
		log.Error().Err(err).Msg("Failed to signUP")
		return fmt.Errorf("failed to sign up: %v", err)
	}

	return nil
}

func (s *AuthService) SignIN(email, password string) (string, error) {
	user, err := s.aRepo.FindByEmail(email)
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
