package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/shiv122/go-todo/app/models"
	"github.com/shiv122/go-todo/connection"
	"github.com/shiv122/go-todo/helpers"
)

type SignUpController struct{}

func (lc *SignUpController) SignUp(c *fiber.Ctx) error {

	hasedPassword, _ := new(helpers.PasswordHelper).HashPassword(c.FormValue("Password"))

	user := models.User{
		FirstName: c.FormValue("FirstName"),
		LastName:  c.FormValue("LastName"),
		Email:     c.FormValue("Email"),
		Password:  hasedPassword,
	}
	result := connection.DB.Create(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": strings.SplitAfter(result.Error.Error(), ":")[1],
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
	})
}
