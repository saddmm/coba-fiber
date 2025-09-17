package helper

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    any `json:"data,omitempty"`
}

func SuccessResponse(c *fiber.Ctx, status int, message string, data any) error {
	return c.Status(status).JSON(Response{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(Response{
		Status:  "error",
		Message: message,
	})
}
func FailResponse(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(Response{
		Status:  "fail",
		Message: message,
	})
}
