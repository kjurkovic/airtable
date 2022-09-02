package service

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/kjurkovic/airtable/service/meta/datastore"
	"github.com/kjurkovic/airtable/service/meta/models"
)

func (service *MetaService) GetOne(rw http.ResponseWriter, r *http.Request) {
	service.Log.Info("Meta service GET one %s", mux.Vars(r)["metaId"])
	id := uuid.MustParse(mux.Vars(r)["metaId"])

	claims := r.Context().Value(models.Claims{}).(*models.Claims)

	meta, err := datastore.MetaDao.GetOne(id)

	if err != nil {
		service.Log.Error(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	// reject request if user is not owner and meta is not publicly available
	if meta.UserId != claims.UserId && !meta.Public {
		service.Log.Info("User not allowed to access meta - not owner, not public")
		rw.WriteHeader(http.StatusForbidden)
		return
	}

	meta.Serialize(rw)
}
