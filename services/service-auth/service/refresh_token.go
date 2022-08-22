package service

import (
	"net/http"

	"github.com/kjurkovic/airtable/service/auth/datastore"
	"github.com/kjurkovic/airtable/service/auth/models"
)

func (service *AuthService) RefreshToken(rw http.ResponseWriter, r *http.Request) {
	service.Log.Info("POST Auth Refresh token")

	request, err := service.validateRefreshTokenRequest(&rw, r)

	if err != nil {
		return
	}

	token, err := datastore.RefreshTokenDao.GetByToken(request.Token)

	if err != nil {
		models.WrongCredentials.SendErrorResponse(rw, http.StatusForbidden)
		return
	}

	storedUser, err := datastore.UserDao.GetById(token.UserId)

	if err != nil {
		models.WrongCredentials.SendErrorResponse(rw, http.StatusForbidden)
		return
	}

	service.generateAuthResponse(storedUser, rw).Serialize(rw)
}

func (service *AuthService) validateRefreshTokenRequest(rw *http.ResponseWriter, r *http.Request) (*models.RefreshTokenRequest, error) {
	request := &models.RefreshTokenRequest{}
	err := request.Deserialize(r.Body)

	if err != nil {
		service.Log.ErrorS("Error deserializing request", r.Body)
		models.SerializationError.SendErrorResponse(*rw, http.StatusBadRequest)
		return nil, err
	}

	return request, nil
}
