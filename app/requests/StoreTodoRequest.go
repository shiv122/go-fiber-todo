package requests

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shiv122/go-todo/helpers"
)

type StoreTodoRequest struct {
	Name        string `form:"Name"  validate:"required,max=500"`
	Description string `form:"Description"  validate:"max=10000"`
}

func (str *StoreTodoRequest) Validate(c *fiber.Ctx) error {
	data := StoreTodoRequest{}
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
