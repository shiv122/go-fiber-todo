package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/shiv122/go-todo/helpers"
)

type UserAuthMiddleware struct{}

func (um *UserAuthMiddleware) Auth(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token format",
		})
	}

	bearerToken := tokenParts[1]

	claims, err := new(helpers.JwtHelper).ExtractClaimsFromToken(bearerToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
			"error":   err,
		})
	}

	c.Locals("user", claims["user"])

	return c.Next()
}
