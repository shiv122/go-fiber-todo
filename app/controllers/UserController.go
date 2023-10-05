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

	var user models.User
	var count int64
	tokenUser := c.Locals("user").(map[string]interface{})
	connection.DB.Where("id = ?", tokenUser["ID"]).First(&user)
	count = connection.DB.Model(&user).Association("Todos").Count()
	return c.JSON(fiber.Map{
		"User":      user,
		"TodoCount": count,
	})
}
