package models

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
	"github.com/kjurkovic/airtable/service/auth/validators"
)

type RegisterRequest struct {
	FirstName string `json:"firstName" validate:"required,name"`
	LastName  string `json:"lastName" validate:"required,name"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,password"`
}

func (request *RegisterRequest) Deserialize(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(request)
}

// validation
func (request *RegisterRequest) Validate() error {
	validator := validator.New()
	validator.RegisterValidation("name", validators.ValidateName)
	validator.RegisterValidation("email", validators.ValidateEmail)
	validator.RegisterValidation("password", validators.ValidatePassword)
	return validator.Struct(request)
}
