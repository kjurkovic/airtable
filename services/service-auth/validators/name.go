package validators

import (
	"regexp"

	"github.com/go-playground/validator"
)

func ValidateName(fieldValidator validator.FieldLevel) bool {
	// requiring User name to be min 4 letters and contain only alphanumeric characters
	regex := regexp.MustCompile(`[a-zA-Z0-9]{4,}`)
	matches := regex.FindAllString(fieldValidator.Field().String(), 1)
	return len(matches) == 1
}
