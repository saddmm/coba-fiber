package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddmm/coba-fiber/internal/service"
)

type PostHandler struct {
	postService *service.PostService
}

func NewPostHandler(postService *service.PostService) *PostHandler {
	return &PostHandler{postService}
}

func (h *PostHandler) CreatePost(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Post created",
	})
}
