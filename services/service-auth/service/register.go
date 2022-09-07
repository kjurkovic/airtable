package service

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/kjurkovic/airtable/service/auth/datastore"
	"github.com/kjurkovic/airtable/service/auth/models"
	"github.com/kjurkovic/airtable/service/auth/util"
	"gorm.io/gorm"
)

func (service *AuthService) Register(rw http.ResponseWriter, r *http.Request) {
	service.Log.Info("POST Auth Register")

	request, err := service.validateRegisterRequest(&rw, r)

	if err != nil {
		return
	}

	user := models.UserFn(
		request.FirstName,
		request.LastName,
		strings.TrimSpace(strings.ToLower(request.Email)),
		util.GeneratePasswordHashWithSalt(request.Password),
		models.CLIENT,
	)

	_, err = datastore.UserDao.Save(user)
	if err != nil {
		service.Log.Error(err)
		models.ServerError.SendErrorResponse(rw, http.StatusInternalServerError)
		return
	}
	response := service.generateAuthResponse(user, rw)
	response.Serialize(rw)
}

func (service *AuthService) validateRegisterRequest(rw *http.ResponseWriter, r *http.Request) (*models.RegisterRequest, error) {
	request := &models.RegisterRequest{}
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

	_, err = datastore.UserDao.GetByEmail(strings.TrimSpace(strings.ToLower(request.Email)))

	if err == gorm.ErrRecordNotFound {
		return request, nil
	}

	models.UserAlreadyExistError.SendErrorResponse(*rw, http.StatusConflict)
	return nil, fmt.Errorf("user already exists: %s", request.Email)
}
