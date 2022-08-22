package models

import (
	"encoding/json"
	"io"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	Id          uuid.UUID      `json:"id" gorm:"primaryKey"`
	UserId      uuid.UUID      `json:"userId" gorm:"index" validate:"required"`
	Type        string         `json:"type" gorm:"index" validate:"required"`
	AuditObject string         `json:"obj" validate:"required"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"` // setting this to turn on soft delete - nothing should be deleted from db related to Audit events
}

func (request *Event) Deserialize(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(request)
}

func (event *Event) Serialize(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(event)
}
