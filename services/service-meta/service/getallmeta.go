package service

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/kjurkovic/airtable/service/meta/datastore"
	"github.com/kjurkovic/airtable/service/meta/util"
)

func (service *MetaService) GetAll(rw http.ResponseWriter, r *http.Request) {
	service.Log.Info("Meta service GET all %s", mux.Vars(r)["id"])
	id := uuid.MustParse(mux.Vars(r)["userId"])

	page, pageSize := util.GetPaginationParams(r)
	content, err := datastore.MetaDao.GetAll(id, page, pageSize)

	if err != nil {
		service.Log.Error(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	content.Serialize(rw)
}
