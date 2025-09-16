package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddmm/coba-fiber/internal/handler"
)

func SetupPostRoutes(router fiber.Router, postHandler *handler.PostHandler) {
	router.Post("/posts", postHandler.CreatePost)
}