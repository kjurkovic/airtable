package models

import (
	"encoding/json"
	"io"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Data struct {
	Id        uuid.UUID      `json:"id" gorm:"primaryKey"`
	MetaId    uuid.UUID      `json:"metaId" gorm:"index" validate:"required"`
	Content   datatypes.JSON `json:"content" gorm:"index" validate:"required"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (request *Data) Deserialize(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(request)
}

func (meta *Data) Serialize(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(meta)
}
