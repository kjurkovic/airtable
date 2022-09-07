package service

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/kjurkovic/airtable/service/auth/datastore"
	"github.com/kjurkovic/airtable/service/auth/middleware"
	"github.com/kjurkovic/airtable/service/auth/models"
	"github.com/kjurkovic/airtable/service/auth/util"
	"github.com/kjurkovic/airtable/service/auth/wrappers"
)

func (service *UserService) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	service.Log.Info("User service DELETE")

	id := uuid.MustParse(mux.Vars(r)["id"])

	user, err := datastore.UserDao.GetById(id)
	affected, err := datastore.UserDao.Delete(id)

	if err != nil {
		httpErr := &models.HttpError{Title: "Error", Message: err.Error()}
		httpErr.SendErrorResponse(rw, http.StatusInternalServerError)
		return
	}

	if affected == 0 {
		models.UserNotFoundError.SendErrorResponse(rw, http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusAccepted)

	auditObj, _ := util.ToJson(user)

	if err != nil {
		auditObj = id.String()
	}

	claims := r.Context().Value(middleware.KeyClaims{}).(*models.Claims)
	wrappers.Audit.SendEvent(claims.UserId, auditObj, wrappers.DeleteUser)
}
