package service

import (
	"net/http"

	"github.com/kjurkovic/airtable/service/meta/datastore"
	"github.com/kjurkovic/airtable/service/meta/models"
)

func (service *MetaService) Create(rw http.ResponseWriter, r *http.Request) {
	service.Log.Info("Meta service POST create")

	meta := &models.Meta{}
	err := meta.Deserialize(r.Body)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	claims := r.Context().Value(models.Claims{}).(*models.Claims)
	meta.UserId = claims.UserId

	result, err := datastore.MetaDao.Create(meta)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, field := range meta.Fields {
		field.MetaId = result.Id
		datastore.FieldDao.Create(&field)
	}

	rw.WriteHeader(http.StatusCreated)
}
