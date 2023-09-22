package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type TodoController struct{}

func (tc *TodoController) GetList(c *fiber.Ctx) error {
	user := c.Locals("user")

	return c.JSON(user)

}
