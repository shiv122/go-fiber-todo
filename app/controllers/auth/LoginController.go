package auth

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/shiv122/go-todo/app/models"
	"github.com/shiv122/go-todo/config"
	"github.com/shiv122/go-todo/connection"
	"github.com/shiv122/go-todo/helpers"
)

type LoginController struct{}

func (lc *LoginController) Login(c *fiber.Ctx) error {

	var user models.User

	result := connection.DB.Where("email = ?", c.FormValue("Email")).First(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).
			JSON(fiber.Map{
				"error":   "404 not found",
				"message": result.Error.Error(),
			})
	}

	passwordMatched := new(helpers.PasswordHelper).CheckPasswordHash(c.FormValue("Password"), user.Password)

	if !passwordMatched {
		return c.Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{
				"error":   "Unauthorized",
				"message": "Email or Password is incorrect",
			})
	}

	claims := jwt.MapClaims{
		"user": map[string]string{
			"ID":        strconv.FormatUint(uint64(user.ID), 10),
			"Email":     user.Email,
			"FirstName": user.FirstName,
			"LastName":  user.LastName,
		},
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.App.SecretKey))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"token":  t,
	})
}
