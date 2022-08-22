package models

import (
	"encoding/json"
	"io"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id        uuid.UUID      `json:"id" gorm:"primaryKey"`
	FirstName string         `json:"firstName" gorm:"index:idx_user_name"`
	LastName  string         `json:"lastName" gorm:"index:idx_user_name"`
	Email     string         `json:"email" gorm:"index"`
	Image     string         `json:"image,omitempty"`
	Password  string         `json:"-"`
	Role      Role           `json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (user *User) Serialize(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(user)
}

func UserFn(firstName string, lastName string, email string, password string, role Role) *User {
	return &User{
		Id:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		Role:      role,
	}
}
