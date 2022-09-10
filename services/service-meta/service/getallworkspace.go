package service

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/kjurkovic/airtable/service/meta/datastore"
	"github.com/kjurkovic/airtable/service/meta/middleware"
	"github.com/kjurkovic/airtable/service/meta/models"
	"github.com/kjurkovic/airtable/service/meta/util"
)

func (service *MetaService) GetAllWorkspace(rw http.ResponseWriter, r *http.Request) {
	service.Log.Info("Meta service GET all by workspace %s", mux.Vars(r)["id"])
	id := uuid.MustParse(mux.Vars(r)["id"])

	claims := r.Context().Value(middleware.KeyClaims{}).(*models.Claims)

	page, pageSize := util.GetPaginationParams(r)
	content, err := datastore.MetaDao.GetAllWorkspace(claims.UserId, id, page, pageSize)

	if err != nil {
		service.Log.Error(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	content.Serialize(rw)
}
