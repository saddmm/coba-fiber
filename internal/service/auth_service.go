package service

import (
	"errors"

	"github.com/saddmm/coba-fiber/internal/model"
	"github.com/saddmm/coba-fiber/internal/repository"
	"github.com/saddmm/coba-fiber/pkg/helper"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) *AuthService {
	return &AuthService{userRepository}
}

func (s *AuthService) Register(user *model.User) error {
	user, err := s.userRepository.FindByEmail(user.Email)
	if err == nil {
		return errors.New("email already exists")
	}
	return s.userRepository.Create(user)
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("password is incorrect")
	}

	token, err := helper.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}