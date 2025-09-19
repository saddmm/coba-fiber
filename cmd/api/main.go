package main

import (
	"log"

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

	db, err := database.ConnectDB(conf.Database)
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}

	// Repository
	userRepository := repository.NewUserRepository(db)
	postRepository := repository.NewPostRepository(db)

	// Service
	userService := service.NewUserService(userRepository)
	postService := service.NewPostService(postRepository)
	authService := service.NewAuthService(userRepository)

	// Handler
	userHandler := handler.NewUserHandler(userService, authService)
	postHandler := handler.NewPostHandler(postService, userService)

	routes.SetupUserRoutes(app, userHandler)
	routes.SetupPostRoutes(app, postHandler)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(conf.Server.Host + ":" + conf.Server.Port)
}
