package middleware

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/saddmm/coba-fiber/pkg/helper"
)

func ValidateDto(dtoType interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := reflect.New(reflect.TypeOf(dtoType)).Interface()

		if err := c.BodyParser(payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := helper.ValidateStruct(payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errors": err,
			})
		}

		c.Locals("validateDTO", payload)

		return c.Next()
	}
}