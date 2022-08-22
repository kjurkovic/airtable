package models

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
	"github.com/kjurkovic/airtable/service/auth/validators"
)

type PasswordChangeRequest struct {
	Password string `json:"password" validate:"required,password"`
}

func (request *PasswordChangeRequest) Deserialize(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(request)
}

// validation
func (request *PasswordChangeRequest) Validate() error {
	validator := validator.New()
	validator.RegisterValidation("password", validators.ValidatePassword)
	return validator.Struct(request)
}
