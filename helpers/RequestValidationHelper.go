package helpers

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Validator struct{}

type IError struct {
	Field   string      `json:"field"`
	Tag     string      `json:"tag"`
	Value   interface{} `json:"value"`
	Message string      `json:"message"`
}

var validate = validator.New()

func (v *Validator) ValidateData(c *fiber.Ctx, data interface{}) []*IError {
	var errors []*IError

	if err := validate.Struct(data); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Value()
			el.Message = strings.SplitAfter(err.Error(), "Error:Field")[1]
			errors = append(errors, &el)
		}
		return errors
	}

	return nil
}
