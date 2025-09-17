package service

import (
	"errors"

	"github.com/saddmm/coba-fiber/internal/model"
	"github.com/saddmm/coba-fiber/internal/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (s *UserService) GetUser() ([]*model.User, error) {
	users, err := s.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	user, err := s.userRepository.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}