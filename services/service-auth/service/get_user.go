package service

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/kjurkovic/airtable/service/auth/datastore"
	"github.com/kjurkovic/airtable/service/auth/models"
)

func (service *UserService) GetUser(rw http.ResponseWriter, r *http.Request) {
	service.Log.Info("User service GET single user")
	id := uuid.MustParse(mux.Vars(r)["id"])

	user, err := datastore.UserDao.GetById(id)

	if err != nil {
		response := &models.HttpError{Title: "Error", Message: err.Error()}
		response.SendErrorResponse(rw, http.StatusBadRequest)
		return
	}

	err = user.Serialize(rw)

	if err != nil {
		models.SerializationResponseError.SendErrorResponse(rw, http.StatusInternalServerError)
		return
	}
}
