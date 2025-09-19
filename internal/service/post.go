package service

import (
	"errors"

	"github.com/saddmm/coba-fiber/internal/model"
	"github.com/saddmm/coba-fiber/internal/repository"
)

type PostService struct {
	postRepository repository.PostRepository
}

func NewPostService(postRepository repository.PostRepository) *PostService {
	return &PostService{postRepository}
}

func (s *PostService) CreatePost(post *model.Post) error {
	return s.postRepository.CreatePost(post)
}

func (s *PostService) GetPostByID(id uint) (*model.Post, error) {
	post, err := s.postRepository.GetPostByID(id)
	if err != nil {
		return nil, errors.New("post not found")
	}
	return post, nil
}