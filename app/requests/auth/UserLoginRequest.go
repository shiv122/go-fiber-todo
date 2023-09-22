package requestAuth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shiv122/go-todo/helpers"
)

type UserLoginRequest struct {
	Email    string `form:"Email"  validate:"required,email"`
	Password string `form:"Password"  validate:"required"`
}

func (ulr *UserLoginRequest) Validate(c *fiber.Ctx) error {
	data := UserLoginRequest{}
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
