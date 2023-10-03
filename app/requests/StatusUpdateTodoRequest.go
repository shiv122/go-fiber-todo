package requests

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shiv122/go-todo/helpers"
)

type StatusUpdateTodoRequest struct {
	Id     string `from:"Id" validate:"required"`
	Status string `from:"Status" validate:"required,oneof=pending completed"`
}

func (str *StatusUpdateTodoRequest) Validate(c *fiber.Ctx) error {
	data := StatusUpdateTodoRequest{}
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	validator := new(helpers.Validator)
	errors := validator.ValidateData(c, data)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	return c.Next()
}
