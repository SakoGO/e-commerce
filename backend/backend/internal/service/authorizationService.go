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
	SignUP(user *model.User) error
	SignIN(user *model.User) error
	FindByUsername(username string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindByPhone(phone string) (*model.User, error)
}

type AuthService struct {
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) FindByUsername(username string) (*model.User, error) {
	return s.repo.FindByUsername(username)
}

func (s *AuthService) FindByEmail(email string) (*model.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *AuthService) FindByPhone(phone string) (*model.User, error) {
	return s.repo.FindByPhone(phone)
}

func (s *AuthService) SignUP(username, email, password, phone string) error {

	existingUser, err := s.repo.FindByUsername(username)
	if err == nil && existingUser != nil {
		return fmt.Errorf("username %s is already taken", username)
	}

	existingEmail, err := s.repo.FindByEmail(email)
	if err == nil && existingEmail != nil {
		return fmt.Errorf("email %s is already taken", email)
	}

	existingPhone, err := s.repo.FindByPhone(phone)
	if err == nil && existingPhone != nil {
		return fmt.Errorf("phone %s is already taken", phone)
	}

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
	if err := s.repo.SignUP(user); err != nil {
		return fmt.Errorf("failed to sign up")
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
