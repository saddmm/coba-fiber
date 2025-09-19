package repository

import (
	"github.com/saddmm/coba-fiber/internal/model"
	"gorm.io/gorm"
)

type PostRepository interface {
	CreatePost(post *model.Post) error
	GetPostByID(id uint) (*model.Post, error)
	GetAllPosts() ([]model.Post, error)
	UpdatePost(post *model.Post) error
	DeletePost(id uint) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db}
}

// CreatePost implements PostRepository.
func (p *postRepository) CreatePost(post *model.Post) error {
	return p.db.Create(&post).Error
}

// DeletePost implements PostRepository.
func (p *postRepository) DeletePost(id uint) error {
	return p.db.Delete(&model.Post{}, id).Error
}

// GetAllPosts implements PostRepository.
func (p *postRepository) GetAllPosts() ([]model.Post, error) {
	var posts []model.Post
	err := p.db.Find(&posts).Error
	return posts, err
}

// GetPostByID implements PostRepository.
func (p *postRepository) GetPostByID(id uint) (*model.Post, error) {
	var post model.Post
	err := p.db.First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

// UpdatePost implements PostRepository.
func (p *postRepository) UpdatePost(post *model.Post) error {
	return p.db.Save(post).Error
}
