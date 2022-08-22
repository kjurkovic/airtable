package service

import (
	"net/http"
)

func (service *UserService) CreateUser(rw http.ResponseWriter, r *http.Request) {
	service.Log.Info("User service CREATE")
	rw.WriteHeader(451)
}
