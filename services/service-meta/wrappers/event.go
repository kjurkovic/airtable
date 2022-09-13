package wrappers

import (
	"encoding/json"
	"io"

	"github.com/google/uuid"
)

type Event struct {
	UserId      uuid.UUID `json:"userId" validate:"required"`
	Type        AuditType `json:"type" validate:"required"`
	AuditObject string    `json:"obj" validate:"required"`
}

func (request *Event) Deserialize(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(request)
}

func (event *Event) Serialize(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(event)
}
