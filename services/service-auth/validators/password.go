package validators

import (
	"regexp"

	"github.com/go-playground/validator"
)

// - 12 characters long
// - Must include capital letters
// - Must include small letters
// - Must include numbers
// - Must include special characters
func ValidatePassword(fieldValidator validator.FieldLevel) bool {
	regex := regexp.MustCompile(`[a-zA-Z0-9!@#$%^&*()_+?<>:";',./\\]{12,}`)
	matches := regex.FindAllString(fieldValidator.Field().String(), 1)
	return len(matches) >= 1
}
