package validators

import (
	"net/mail"

	"github.com/go-playground/validator"
)

func ValidateEmail(fieldValidator validator.FieldLevel) bool {
	_, err := mail.ParseAddress(fieldValidator.Field().String())
	return err == nil
}
