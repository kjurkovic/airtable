package models

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
	"github.com/kjurkovic/airtable/service/auth/validators"
)

type UserEditRequest struct {
	FirstName string `json:"firstName" validate:"required,name"`
	LastName  string `json:"lastName" validate:"required,name"`
	Email     string `json:"email" validate:"required,email"`
}

func (request *UserEditRequest) Deserialize(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(request)
}

// validation
func (request *UserEditRequest) Validate() error {
	validator := validator.New()
	validator.RegisterValidation("name", validators.ValidateName)
	validator.RegisterValidation("email", validators.ValidateEmail)
	return validator.Struct(request)
}
