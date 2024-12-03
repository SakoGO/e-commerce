package service

import (
	"e-commerce/backend/internal/model"
)

//TODO: Функции юзера:

type UserRepository interface {
	UserFindByID(userID int) (*model.User, error) //
	//UserDelete(userID int) error                  //
}

type UserService struct {
	repo UserRepository
}

func (s *UserService) UserFindByID(userID int) (*model.User, error) {
	return s.repo.UserFindByID(userID)
}

//func (s *UserService) UserDelete(userID int) error {
//	return s.repo.UserDelete(userID)
//}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}
