package service

import (
	"net/http"
	"strings"

	"github.com/kjurkovic/airtable/service/auth/datastore"
	"github.com/kjurkovic/airtable/service/auth/models"
	"github.com/kjurkovic/airtable/service/auth/util"
)

func (service *AuthService) Login(rw http.ResponseWriter, r *http.Request) {
	service.Log.Info("POST Auth Login")

	request, err := service.validateLoginRequest(&rw, r)

	if err != nil {
		return
	}

	storedUser, err := datastore.UserDao.GetByEmail(strings.TrimSpace(strings.ToLower(request.Email)))

	if err != nil {
		models.WrongCredentials.SendErrorResponse(rw, http.StatusForbidden)
		return
	}

	isUserAuthorized := service.validatePassword(storedUser.Password, request.Password)

	if !isUserAuthorized {
		models.WrongCredentials.SendErrorResponse(rw, http.StatusForbidden)
		return
	}

	service.generateAuthResponse(storedUser, rw).Serialize(rw)
}

func (service *AuthService) validateLoginRequest(rw *http.ResponseWriter, r *http.Request) (*models.LoginRequest, error) {
	request := &models.LoginRequest{}
	err := request.Deserialize(r.Body)

	if err != nil {
		service.Log.ErrorS("Error deserializing request", r.Body)
		models.SerializationError.SendErrorResponse(*rw, http.StatusBadRequest)
		return nil, err
	}

	err = request.Validate()

	if err != nil {
		httpErr := models.UserValidationError
		httpErr.Message = err.Error()
		httpErr.SendErrorResponse(*rw, http.StatusBadRequest)
		return nil, err
	}

	return request, nil
}

func (service *AuthService) validatePassword(storedPassword string, inputPassword string) bool {
	components := strings.Split(storedPassword, ":")
	salt := components[0]
	encryptedPassword := components[1]

	requestPasswordHash := util.CreateHash(salt, inputPassword)
	return encryptedPassword == requestPasswordHash
}
