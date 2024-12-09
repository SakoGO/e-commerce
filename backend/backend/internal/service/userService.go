package service

import (
	"e-commerce/backend/internal/model"
	"errors"
)

//TODO: Функции юзера:

type UserRepository interface {
	UserFindByID(userID int) (*model.User, error) //
	UserUpdate(user *model.User) error
	//UserDelete(userID int) error                  //
}

type UserService struct {
	repo UserRepository
}

func (s *UserService) UserFindByID(userID int) (*model.User, error) {
	return s.repo.UserFindByID(userID)
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) UserUpdate(userID int, user *model.User) (*model.User, error) {
	existingUser, err := s.repo.UserFindByID(userID)
	if err != nil {
		return nil, err
	}
	existingUser.Username = user.Username
	existingUser.Email = user.Email
	existingUser.Password = user.Password
	existingUser.Phone = user.Phone

	err = s.repo.UserUpdate(existingUser)
	if err != nil {
		return nil, errors.New("failed to update user") // Если произошла ошибка при сохранении
	}
	return existingUser, nil
}
