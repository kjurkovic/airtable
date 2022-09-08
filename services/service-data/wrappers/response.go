package wrappers

import (
	"encoding/json"
	"io"

	"github.com/google/uuid"
)

type Field struct {
	Id         uuid.UUID `json:"id""`
	MetaId     uuid.UUID `json:"metaId"`
	Label      string    `json:"label"`
	Type       string    `json:"type"`
	Validation string    `json:"validation"`
}

type Meta struct {
	Id          uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"userId"`
	WorkspaceId uuid.UUID `json:"workspaceId"`
	Fields      []Field   `json:"fields"`
	Name        string    `json:"name"`
	Public      bool      `json:"isPublic"`
}

func (request *Meta) Deserialize(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(request)
}

func (meta *Meta) Serialize(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(meta)
}

func (field *Field) Deserialize(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(field)
}

func (field *Field) Serialize(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(field)
}
