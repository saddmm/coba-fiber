package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddmm/coba-fiber/internal/handler"
	"github.com/saddmm/coba-fiber/internal/repository"
	"github.com/saddmm/coba-fiber/internal/routes"
	"github.com/saddmm/coba-fiber/internal/service"
	"github.com/saddmm/coba-fiber/pkg/config"
	"github.com/saddmm/coba-fiber/pkg/database"
)

func main() {
	app := fiber.New()

	conf := config.Get()

	database.ConnectDB(conf.Database)

	// Repository
	userRepository := repository.NewUserRepository(database.DB)
	postRepository := repository.NewPostRepository(database.DB)

	// Service
	userService := service.NewUserService(userRepository)
	postService := service.NewPostService(postRepository)
	authService := service.NewAuthService(userRepository)

	// Handler
	userHandler := handler.NewUserHandler(userService, authService)
	postHandler := handler.NewPostHandler(postService)

	routes.SetupUserRoutes(app, userHandler)
	routes.SetupPostRoutes(app, postHandler)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(conf.Server.Host + ":" + conf.Server.Port)
}
