package service

import (
	"github.com/saddmm/coba-fiber/internal/model"
	"github.com/saddmm/coba-fiber/internal/repository"
)

type PostService struct {
	postRepository repository.PostRepository
}

func NewPostService(postRepository repository.PostRepository) *PostService {
	return &PostService{postRepository}
}

func (s *PostService) CreatePost(userId uint, post *model.Post) error {
	return s.postRepository.CreatePost(userId, post)
}
