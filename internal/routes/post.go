package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddmm/coba-fiber/internal/dto"
	"github.com/saddmm/coba-fiber/internal/handler"
	"github.com/saddmm/coba-fiber/internal/middleware"
)

func SetupPostRoutes(router fiber.Router, postHandler *handler.PostHandler) {
	router.Post("/posts", middleware.Protected(), middleware.ValidateDto(dto.CreatePostDto{}), postHandler.CreatePost)
	router.Get("/posts/:id", postHandler.GetPost)
}