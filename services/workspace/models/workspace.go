package models

import (
	"encoding/json"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

// Workspace model
// swagger:model Workspace
type Workspace struct {
	// Workspace ID
	//
	// required: true
	ID uuid.UUID `json:"id" gorm:"primaryKey"`
	// Workspace name
	//
	// required: true
	// minLength: 4
	Name      string    `json:"name" validate:"required,name"`
	UserId    uuid.UUID `json:"-" gorm:"index"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Workspaces []*Workspace

// serialization
func (workspace *Workspaces) Serialize(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(workspace)
}

func (workspace *Workspace) Serialize(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(workspace)
}

func (workspace *Workspace) Deserialize(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(workspace)
}

// validation
func (workspace *Workspace) Validate() error {
	validator := validator.New()
	validator.RegisterValidation("name", validateName)
	return validator.Struct(workspace)
}

func validateName(fieldValidator validator.FieldLevel) bool {
	// requiring Workspace name to be min 4 letters and contain only alphanumeric characters
	regex := regexp.MustCompile(`[a-zA-Z0-9]{4,}`)
	matches := regex.FindAllString(fieldValidator.Field().String(), 1)

	return len(matches) == 1
}
