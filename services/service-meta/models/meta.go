package models

import (
	"encoding/json"
	"io"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Type string

const (
	Form Type = "form"
)

type Field struct {
	Id        uuid.UUID      `json:"id" gorm:"primaryKey"`
	MetaId    uuid.UUID      `json:"metaId" gorm:"index"`
	Type      Type           `json:"type"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Meta struct {
	Id          uuid.UUID      `json:"id" gorm:"primaryKey"`
	UserId      uuid.UUID      `json:"userId" gorm:"index" validate:"required"`
	WorkspaceId uuid.UUID      `json:"workspaceId" gorm:"index" validate:"required"`
	Fields      []Field        `json:"fields" gorm:"index" validate:"required"`
	Name        string         `json:"name" validate:"required"`
	Public      bool           `json:"is_public" gorm:"index"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
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
