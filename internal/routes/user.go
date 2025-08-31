package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddmm/coba-fiber/internal/handler"
)

func SetupUserRoutes(router fiber.Router, userHandler *handler.UserHandler) {
	userRoutes := router.Group("/users")

	userRoutes.Post("/", userHandler.CreateUser)
}