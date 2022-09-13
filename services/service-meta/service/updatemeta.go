package service

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/kjurkovic/airtable/service/meta/datastore"
	"github.com/kjurkovic/airtable/service/meta/middleware"
	"github.com/kjurkovic/airtable/service/meta/models"
	"github.com/kjurkovic/airtable/service/meta/util"
	"github.com/kjurkovic/airtable/service/meta/wrappers"
)

func (service *MetaService) Update(rw http.ResponseWriter, r *http.Request) {
	service.Log.Info("Meta service POST create")

	id := uuid.MustParse(mux.Vars(r)["id"])

	claims := r.Context().Value(middleware.KeyClaims{}).(*models.Claims)

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

	auditObj, _ := util.ToJson(model)
	wrappers.Audit.SendEvent(claims.UserId, auditObj, wrappers.MetaCreated)
}
