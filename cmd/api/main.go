package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddmm/coba-fiber/internal/handler"
	"github.com/saddmm/coba-fiber/internal/model"
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
	database.DB.AutoMigrate(&model.User{}, &model.Post{})

	// Repository
	userRepository := repository.NewUserRepository(database.DB)
	postRepository := repository.NewPostRepository(database.DB)

	// Service
	userService := service.NewUserService(userRepository)
	postService := service.NewPostService(postRepository)

	// Handler
	userHandler := handler.NewUserHandler(userService)
	postHandler := handler.NewPostHandler(postService)

	routes.SetupUserRoutes(app, userHandler)
	routes.SetupPostRoutes(app, postHandler)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(conf.Server.Host + ":" + conf.Server.Port)
}
