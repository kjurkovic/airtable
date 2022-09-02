package models

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type Claims struct {
	UserId      uuid.UUID `json:"userId"`
	Username    string    `json:"username"`
	ProductRole Role      `json:"role"`
	jwt.StandardClaims
}
