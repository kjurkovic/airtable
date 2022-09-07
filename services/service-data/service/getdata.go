package service

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/kjurkovic/airtable/service/data/datastore"
	"github.com/kjurkovic/airtable/service/data/util"
)

func (service *DataService) Get(rw http.ResponseWriter, r *http.Request) {
	service.Log.Info("Data service GET for meta id %s", mux.Vars(r)["metaId"])

	id := uuid.MustParse(mux.Vars(r)["metaId"])
	page, pageSize := util.GetPaginationParams(r)

	content, err := datastore.DataDao.GetAll(id, page, pageSize)

	if err != nil {
		service.Log.Error(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	content.Serialize(rw)
}
