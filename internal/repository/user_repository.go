package repository

import (
	"github.com/saddmm/coba-fiber/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByID(id uint) (*model.User, error)
	FindAll() ([]*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Update(user *model.User) (model.User, error)
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

// Create implements UserRepository.
func (u *userRepository) Create(user *model.User) error {
	return u.db.Create(user).Error
}

// Delete implements UserRepository.
func (u *userRepository) Delete(id uint) error {
	return u.db.Delete(&model.User{}, id).Error
}

// FindAll implements UserRepository.
func (u *userRepository) FindAll() ([]*model.User, error) {
	var users []*model.User
	err := u.db.Find(&users).Error
	return users, err
}

// FindByID implements UserRepository.
func (u *userRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := u.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update implements UserRepository.
func (u *userRepository) Update(user *model.User) (model.User, error) {
	err := u.db.Save(user).Error
	return *user, err
}

func (u *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

