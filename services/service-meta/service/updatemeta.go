package service

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/kjurkovic/airtable/service/meta/datastore"
	"github.com/kjurkovic/airtable/service/meta/models"
)

func (service *MetaService) Update(rw http.ResponseWriter, r *http.Request) {
	service.Log.Info("Meta service POST create")

	id := uuid.MustParse(mux.Vars(r)["id"])

	model := &models.Meta{}
	err := model.Deserialize(r.Body)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	model.Id = id

	err = datastore.MetaDao.Update(*model)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusAccepted)
}
