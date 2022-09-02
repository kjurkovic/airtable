package service

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/kjurkovic/airtable/service/meta/datastore"
	"github.com/kjurkovic/airtable/service/meta/models"
)

func (service *MetaService) Delete(rw http.ResponseWriter, r *http.Request) {
	service.Log.Info("Meta service DELETE %s", mux.Vars(r)["id"])

	id := uuid.MustParse(mux.Vars(r)["id"])
	claims := r.Context().Value(models.Claims{}).(*models.Claims)

	err := datastore.MetaDao.Delete(id, claims.UserId)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusAccepted)
}
