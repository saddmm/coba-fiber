package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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

func (h *UserHandler) Login(c *fiber.Ctx) error {
	input := c.Locals("validateDTO").(*dto.LoginDto)
	
	token, err := h.authService.Login(input.Email, input.Password)
	if err != nil {
		if err.Error() == "user not found" || err.Error() == "password is incorrect" {
			return helper.ErrorResponse(c, fiber.StatusUnauthorized, err.Error())
		}
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return helper.SuccessResponse(c, fiber.StatusOK, "Login successful", fiber.Map{"token": token})
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	users, err := h.userService.GetUser()
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return helper.SuccessResponse(c, fiber.StatusOK, "Users retrieved successfully", users)
}

func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
	userToken := c.Locals("user").(*jwt.Token)
	user := userToken.Claims.(*helper.TokenClaims)
	if user == nil {
		return helper.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid token")
	}

	return helper.SuccessResponse(c, fiber.StatusOK, "User profile fetched successfully", user)
}