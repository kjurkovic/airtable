package service

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/kjurkovic/airtable/service/meta/datastore"
	"github.com/kjurkovic/airtable/service/meta/middleware"
	"github.com/kjurkovic/airtable/service/meta/models"
	"github.com/kjurkovic/airtable/service/meta/wrappers"
)

func (service *MetaService) Delete(rw http.ResponseWriter, r *http.Request) {
	service.Log.Info("Meta service DELETE %s", mux.Vars(r)["id"])

	id := uuid.MustParse(mux.Vars(r)["id"])
	claims := r.Context().Value(middleware.KeyClaims{}).(*models.Claims)

	err := datastore.MetaDao.Delete(id, claims.UserId)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusAccepted)

	wrappers.Audit.SendEvent(claims.UserId, id.String(), wrappers.MetaDeleted)
}
