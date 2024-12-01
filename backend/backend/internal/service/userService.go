package service

import (
	"e-commerce/backend/internal/model"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type UserRepository interface {
	UserSignUP(user *model.User) error
	UserFindByUsername(username string) (*model.User, error)
	UserFindByEmail(email string) (*model.User, error)
	//	UserFindByID(userID int) (*model.User, error)
	//	UserDelete(userID int) error
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) SignUP(username, email, password string) error {
	existingUser, err := s.repo.UserFindByUsername(username)
	if err == nil && existingUser != nil {
		return fmt.Errorf("username %s is already taken", username)
	}

	existingEmail, err := s.repo.UserFindByEmail(email)
	if err == nil && existingEmail != nil {
		return fmt.Errorf("email %s is already taken", email)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("failed to hash password")
		return err
	}

	user := &model.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}
	if err := s.repo.UserSignUP(user); err != nil {
		return fmt.Errorf("failed to sign up")
	}
	return nil
}

func (s *UserService) SignIN(email, password string) (string, error) {
	user, err := s.repo.UserFindByEmail(email)
	if err != nil {
		return "", fmt.Errorf("account %s is not found", email)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("incorrect passwrod")
	}

	token, err := s.GenerateJWTToken(user.UserID)
	if err != nil {
		log.Error().Err(err).Str("userID", fmt.Sprintf("%d", user.UserID)).Msg("Error generating token")
		return "", err
	}
	return token, nil
}

func (s *UserService) GenerateJWTToken(userID int) (string, error) {
	keyJWT := s.GetJWTKey()

	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 17).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(keyJWT))
	if err != nil {
		log.Error().Msg("Error creating token")
		return "", fmt.Errorf("unable to create token: %v", err)
	}
	return tokenString, nil
}

func (s *UserService) GetJWTKey() string {
	keyJWT := os.Getenv("JWT_SECRET_KEY")
	fmt.Println(len(keyJWT))
	if keyJWT == "" {
		log.Fatal().Msg("JWT secret key is not set")
	}
	return keyJWT
}
