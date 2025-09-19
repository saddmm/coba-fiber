package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/saddmm/coba-fiber/internal/dto"
	"github.com/saddmm/coba-fiber/internal/model"
	"github.com/saddmm/coba-fiber/internal/service"
	"github.com/saddmm/coba-fiber/pkg/helper"
)

type PostHandler struct {
	postService *service.PostService
	userService *service.UserService
}

func NewPostHandler(postService *service.PostService, userService *service.UserService) *PostHandler {
	return &PostHandler{postService, userService}
}

func (h *PostHandler) CreatePost(c *fiber.Ctx) error {
	userID := c.Locals("user").(*jwt.Token).Claims.(*helper.TokenClaims).UserID
	
	input := c.Locals("validateDTO").(*dto.CreatePostDto)
	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusNotFound, err.Error())
	}

	post := model.Post{
		Title:   input.Title,
		Content: input.Content,
		UserID:  userID,
		User: *user,
	}

	if err := h.postService.CreatePost(&post); err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return helper.SuccessResponse(c, fiber.StatusCreated, "Post created successfully", post)
}

func (h *PostHandler) GetPost(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Invalid post ID")
	}

	post, err := h.postService.GetPostByID(uint(id))
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusNotFound, err.Error())
	}
	return helper.SuccessResponse(c, fiber.StatusOK, "Post retrieved successfully", post)
}