package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/kjurkovic/airtable/service/data/datastore"
	"github.com/kjurkovic/airtable/service/data/models"
	"github.com/kjurkovic/airtable/service/data/wrappers"
)

func (service *DataService) Create(rw http.ResponseWriter, r *http.Request) {
	service.Log.Info("Data service POST for metaId %s", mux.Vars(r)["metaId"])
	id := uuid.MustParse(mux.Vars(r)["metaId"])

	model := &models.Data{}
	model.Deserialize(r.Body)
	model.MetaId = id
	model.Id = uuid.New()

	_, err := wrappers.MetaApi.Get(id)

	if err != nil {
		service.Log.Error(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	var content map[string]interface{}
	err = json.Unmarshal([]byte(model.Content.String()), &content)

	if err != nil {
		service.Log.Error(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	meta, err := wrappers.MetaApi.Get(id)

	if err != nil {
		service.Log.Error(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	validated := service.validateContent(content, meta)

	if !validated {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = datastore.DataDao.Create(model)

	if err != nil {
		service.Log.Error(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusAccepted)
}

func (service *DataService) validateContent(content map[string]interface{}, meta *wrappers.Meta) bool {

	result := true

	for key, element := range content {
		if innerContent, ok := element.(map[string]interface{}); ok {
			result = result && service.validateContent(innerContent, meta)
		} else {
			result = result && service.validateDataField(key, element, meta)
		}
	}
	return result
}

func (service *DataService) validateDataField(key string, element interface{}, meta *wrappers.Meta) bool {

	result := false

	for _, metaField := range meta.Fields {
		if metaField.Label == key {
			fmt.Println(metaField.Validation)
			if metaField.Validation != "" {
				match, _ := regexp.MatchString(metaField.Validation, element.(string))
				result = match
			} else {
				result = true
			}
			break
		}
	}
	return result
}
