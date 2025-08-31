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
	database.DB.AutoMigrate(&model.User{})

	userRepository := repository.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(*userService)

	routes.SetupUserRoutes(app, userHandler)

	app.Use()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(conf.Server.Host + ":" + conf.Server.Port)
}
