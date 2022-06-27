package models

import (
	"encoding/json"
	"io"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

// User model
// swagger:model User
type User struct {
	// User ID
	ID uuid.UUID `json:"id" gorm:"primaryKey"`
	// User name
	//
	// minLength: 4
	Name string `json:"name" validate:"name"`
	// User email
	//
	// required: true
	Email string `json:"email" validate:"required,email"`
	// User password
	//
	// required: true
	Password            string    `json:"password,omitempty" validate:"required,password"`
	RefreshToken        string    `json:"-"`
	ForgotPasswordToken string    `json:"-"`
	CreatedAt           time.Time `json:"-"`
	UpdatedAt           time.Time `json:"-"`
}

// serialization
func (user *User) Serialize(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(user)
}

func (user *User) Deserialize(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(user)
}

// validation
func (user *User) ValidateRegister() error {
	validator := validator.New()
	validator.RegisterValidation("name", validateNameRegistration)
	validator.RegisterValidation("email", validateEmail)
	validator.RegisterValidation("password", validatePassword)
	return validator.Struct(user)
}

func (user *User) ValidateLogin() error {
	validator := validator.New()
	validator.RegisterValidation("name", validateNameLogin)
	validator.RegisterValidation("email", validateEmail)
	validator.RegisterValidation("password", validatePassword)
	return validator.Struct(user)
}
