package models

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

type ResetPasswordRequest struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

type Claims struct {
	UserId      uuid.UUID `json:"userId"`
	Username    string    `json:"username"`
	ProductRole Role      `json:"role"`
	jwt.StandardClaims
}
