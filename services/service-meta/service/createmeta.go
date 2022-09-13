package service

import (
	"net/http"

	"github.com/kjurkovic/airtable/service/meta/datastore"
	"github.com/kjurkovic/airtable/service/meta/middleware"
	"github.com/kjurkovic/airtable/service/meta/models"
	"github.com/kjurkovic/airtable/service/meta/util"
	"github.com/kjurkovic/airtable/service/meta/wrappers"
)

func (service *MetaService) Create(rw http.ResponseWriter, r *http.Request) {
	service.Log.Info("Meta service POST create")

	meta := &models.Meta{}
	err := meta.Deserialize(r.Body)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	claims := r.Context().Value(middleware.KeyClaims{}).(*models.Claims)
	meta.UserId = claims.UserId

	_, err = datastore.MetaDao.Create(meta)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusCreated)

	auditObj, _ := util.ToJson(meta)
	wrappers.Audit.SendEvent(claims.UserId, auditObj, wrappers.MetaCreated)
}
