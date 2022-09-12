package models

import (
	"encoding/json"
	"io"
)

type Validations []*Validation

type Validation struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

// serialization
func (validations *Validations) Serialize(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(validations)
}
