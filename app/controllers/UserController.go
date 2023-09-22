package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shiv122/go-todo/app/models"
	"github.com/shiv122/go-todo/connection"
)

type UserController struct{}

func (uc *UserController) GetUsers(c *fiber.Ctx) error {

	var users []models.User

	result := connection.DB.Find(&users)

	if result.Error != nil {
		return result.Error
	}
	res := map[string]interface{}{
		"users": users,
	}
	return c.JSON(res)
}

func (uc *UserController) GetProfile(c *fiber.Ctx) error {

	return c.JSON(c.Locals("user"))
}
