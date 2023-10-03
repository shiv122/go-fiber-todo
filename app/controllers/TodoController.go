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
	connection.DB.Find(&todos, "`user_id` = ?", user["ID"])

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

func (tc *TodoController) UpdateTodo(c *fiber.Ctx) error {
	user := c.Locals("user").(map[string]interface{})
	var todo models.Todo

	connection.DB.Model(&todo).
		Where("`user_id`=?", user["ID"]).
		Where("id", c.FormValue("Id")).
		Update("name", c.FormValue("Name")).
		Update("description", c.FormValue("Description"))

	return c.JSON(fiber.Map{
		"message": "Todo Updated",
		"status":  "success",
	})

}

func (tc *TodoController) DeleteTodo(c *fiber.Ctx) error {
	user := c.Locals("user").(map[string]interface{})
	todo := models.Todo{}

	connection.DB.Unscoped().
		Where("`user_id`=?", user["ID"]).
		Delete(&todo, c.Params("id"))

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Todo Deleted Successfully",
	})
}

func (tc *TodoController) UpdateStatus(c *fiber.Ctx) error {
	user := c.Locals("user").(map[string]interface{})
	var todo models.Todo

	connection.DB.Model(&todo).
		Where("`user_id`=?", user["ID"]).
		Where("id", c.FormValue("Id")).
		Update("status", c.FormValue("Status"))

	return c.JSON("success")

}
