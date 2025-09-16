package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddmm/coba-fiber/internal/dto"
	"github.com/saddmm/coba-fiber/internal/handler"
	"github.com/saddmm/coba-fiber/internal/middleware"
)

func SetupUserRoutes(router fiber.Router, userHandler *handler.UserHandler) {
	userRoutes := router.Group("/users")

	userRoutes.Post("/", middleware.ValidateDto(dto.CreateUserDto{}), userHandler.CreateUser)
}