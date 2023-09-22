package requestAuth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shiv122/go-todo/helpers"
)

type UserSignUpRequest struct {
	FirstName string `form:"FirstName"  validate:"required"`
	LastName  string `form:"LastName"  validate:"required"`
	Email     string `form:"Email"  validate:"required,email"`
	Password  string `form:"Password"  validate:"required,min=8"`
	Phone     string `from:"Phone" validate:"isdefault"`
}

func (usv *UserSignUpRequest) Validate(c *fiber.Ctx) error {
	data := UserSignUpRequest{}
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
