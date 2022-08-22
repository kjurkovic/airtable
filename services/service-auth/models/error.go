package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpError struct {
	Title   string `json:"title"`
	Message string `json:"description"`
}

var (
	SerializationError         = &HttpError{Title: "Parsing error", Message: "Unable to parse request body. Please check sent data."}
	SerializationResponseError = &HttpError{Title: "Parsing error", Message: "Unable to create response body. Please contact support."}
	UserAlreadyExistError      = &HttpError{Title: "User already exists", Message: "User with this email already exists."}
	UserValidationError        = &HttpError{Title: "User validation error", Message: "Something went wrong with validating your data. Check your name and email."}
	PasswordValidationError    = &HttpError{Title: "User validation error", Message: "Password has to have 12 characters, 1 uppercase, 1 lowercase and 1 special character."}
	UserNotFoundError          = &HttpError{Title: "User not found", Message: "User you are trying to update does not exist."}
	WrongCredentials           = &HttpError{Title: "Login error", Message: "Check your credentials and try again."}
	ServerError                = &HttpError{Title: "Server error", Message: "Something went wrong. Please try again."}
)

func (err *HttpError) SendErrorResponse(rw http.ResponseWriter, code int) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("X-Content-Type-Options", "nosniff")
	rw.WriteHeader(code)
	encoder := json.NewEncoder(rw)
	fmt.Fprint(rw, encoder.Encode(err))
}
