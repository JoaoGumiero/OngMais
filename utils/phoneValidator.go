package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func init() {
	validate := validator.New()
	validate.RegisterValidation("phone_br", validateBRPhone)
}

func validateBRPhone(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^\(?[1-9]{2}\)?\s?9?[0-9]{4}-?[0-9]{4}$`)
	return re.MatchString(fl.Field().String())
}
