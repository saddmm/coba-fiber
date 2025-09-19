package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddmm/coba-fiber/internal/dto"
	"github.com/saddmm/coba-fiber/internal/model"
	"github.com/saddmm/coba-fiber/internal/service"
	"github.com/saddmm/coba-fiber/pkg/helper"
)

type UserHandler struct {
	userService *service.UserService
	authService  *service.AuthService
}

func NewUserHandler(userService *service.UserService, authService *service.AuthService) *UserHandler {
	return &UserHandler{userService, authService}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	input := c.Locals("validateDTO").(*dto.RegisterDto)

	user := model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	
	if err := h.authService.Register(&user); err != nil {
		if err.Error() == "email already exists" {
			return helper.ErrorResponse(c, fiber.StatusConflict, err.Error())
		}
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return helper.SuccessResponse(c, fiber.StatusCreated, "User registered successfully", user)
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	users, err := h.userService.GetUser()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(users)
}