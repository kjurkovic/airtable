package models

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
	"github.com/kjurkovic/airtable/service/auth/validators"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (request *LoginRequest) Deserialize(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(request)
}

// validation
func (request *LoginRequest) Validate() error {
	validator := validator.New()
	validator.RegisterValidation("email", validators.ValidateEmail)
	return validator.Struct(request)
}
