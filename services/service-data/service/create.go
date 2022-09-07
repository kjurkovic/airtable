package service

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/kjurkovic/airtable/service/data/datastore"
	"github.com/kjurkovic/airtable/service/data/models"
)

func (service *DataService) Create(rw http.ResponseWriter, r *http.Request) {
	service.Log.Info("Data service POST for metaId %s", mux.Vars(r)["metaId"])
	id := uuid.MustParse(mux.Vars(r)["metaId"])

	model := &models.Data{}
	model.Deserialize(r.Body)
	model.MetaId = id
	model.Id = uuid.New()

	//TODO get meta object by id

	//TODO validate fields based on field validations

	_, err := datastore.DataDao.Create(model)

	if err != nil {
		service.Log.Error(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusAccepted)
}
