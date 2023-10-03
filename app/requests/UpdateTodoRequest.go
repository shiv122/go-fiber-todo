package requests

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shiv122/go-todo/helpers"
)

type UpdateTodoRequest struct {
	Name        string `form:"Name"  validate:"required,max=500"`
	Description string `form:"Description"  validate:"required,max=10000"`
	Id          string `from:"Id" validate:"required"`
}

func (str *UpdateTodoRequest) Validate(c *fiber.Ctx) error {
	data := UpdateTodoRequest{}
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
