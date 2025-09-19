package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddmm/coba-fiber/internal/dto"
	"github.com/saddmm/coba-fiber/internal/handler"
	"github.com/saddmm/coba-fiber/internal/middleware"
)

func SetupUserRoutes(router fiber.Router, userHandler *handler.UserHandler) {
	userRoutes := router.Group("/users")

	userRoutes.Post("/register", middleware.ValidateDto(dto.RegisterDto{}), middleware.Protected(), userHandler.Register)
	userRoutes.Post("/login", middleware.ValidateDto(dto.LoginDto{}), userHandler.Login)
	userRoutes.Get("/profile", middleware.Protected(), userHandler.GetProfile)
	userRoutes.Get("/", userHandler.GetUser)
}