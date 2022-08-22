package service

import (
	"net/http"

	"github.com/kjurkovic/airtable/service/auth/datastore"
	"github.com/kjurkovic/airtable/service/auth/models"
	"github.com/kjurkovic/airtable/service/auth/util"
)

func (service *UserService) GetUsers(rw http.ResponseWriter, r *http.Request) {
	service.Log.Info("User service GET all users")

	page, pageSize := util.GetPaginationParams(r)

	users, err := datastore.UserDao.GetAll(page, pageSize)

	if err != nil {
		response := &models.HttpError{Title: "Error", Message: err.Error()}
		response.SendErrorResponse(rw, http.StatusBadRequest)
		return
	}

	err = users.Serialize(rw)

	if err != nil {
		models.SerializationResponseError.SendErrorResponse(rw, http.StatusInternalServerError)
		return
	}
}
