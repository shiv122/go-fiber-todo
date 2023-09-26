package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/shiv122/go-todo/app/models"
	"github.com/shiv122/go-todo/connection"
)

type TodoController struct{}

func (tc *TodoController) GetList(c *fiber.Ctx) error {
	user := c.Locals("user").(map[string]interface{})
	var todos []models.Todo
	connection.DB.First(&todos, user["ID"])
	return c.JSON(todos)

}

func (tc *TodoController) StoreTodo(c *fiber.Ctx) error {
	user := c.Locals("user").(map[string]interface{})
	userIDStr := user["ID"].(string)
	userID, _ := strconv.ParseUint(userIDStr, 10, 64)

	todo := models.Todo{
		Name:        c.FormValue("Name"),
		Description: c.FormValue("Description"),
		UserID:      uint(userID),
	}

	res := connection.DB.Create(&todo)

	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": res.Error.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"todo":   todo,
	})
}
