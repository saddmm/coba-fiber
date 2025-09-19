package middleware

import (
	"os"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/saddmm/coba-fiber/pkg/helper"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{
			Key: []byte(os.Getenv("JWT_SECRET")),
		},
		Claims: &helper.TokenClaims{},
		ContextKey: "user",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return helper.ErrorResponse(c, fiber.StatusUnauthorized, err.Error())
		},
	})
}