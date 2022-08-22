package service

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/kjurkovic/airtable/service/auth/datastore"
	"github.com/kjurkovic/airtable/service/auth/models"
)

func (service *UserService) EditUser(rw http.ResponseWriter, r *http.Request) {
	id := uuid.MustParse(mux.Vars(r)["id"])
	request, err := service.validateEditRequest(&rw, r, id)

	if err != nil {
		return
	}

	user, err := datastore.UserDao.GetById(id)

	if err != nil {
		models.UserNotFoundError.SendErrorResponse(rw, http.StatusBadRequest)
		return
	}

	err = datastore.UserDao.Update(&models.User{
		Id:        id,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
	})

	if err != nil {
		response := &models.HttpError{
			Title:   "Error",
			Message: err.Error(),
		}
		response.SendErrorResponse(rw, http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusAccepted)
}

func (service *UserService) validateEditRequest(rw *http.ResponseWriter, r *http.Request, id uuid.UUID) (*models.UserEditRequest, error) {
	request := &models.UserEditRequest{}
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

	existingUser, err := datastore.UserDao.GetByEmail(request.Email)

	if err == nil && existingUser.Id != id {
		models.UserAlreadyExistError.SendErrorResponse(*rw, http.StatusConflict)
		return nil, fmt.Errorf("user with email %s already exists", request.Email)
	}

	return request, nil
}
