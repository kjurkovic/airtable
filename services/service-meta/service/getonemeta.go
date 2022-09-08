package service

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/kjurkovic/airtable/service/meta/datastore"
)

func (service *MetaService) GetOne(rw http.ResponseWriter, r *http.Request) {
	service.Log.Info("Meta service GET one %s", mux.Vars(r)["metaId"])
	id := uuid.MustParse(mux.Vars(r)["metaId"])

	meta, err := datastore.MetaDao.GetOne(id)

	if err != nil {
		service.Log.Error(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	meta.Serialize(rw)
}
