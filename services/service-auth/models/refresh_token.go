package models

import (
	"encoding/json"
	"io"

	"github.com/google/uuid"
)

type RefreshToken struct {
	Id        uuid.UUID `gorm:"primaryKey"`
	UserId    uuid.UUID
	Token     string `gorm:"index"`
	ExpiresAt int64
}

type RefreshTokenRequest struct {
	Token string `json:"token" validate:"required"`
}

func (request *RefreshTokenRequest) Deserialize(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(request)
}
