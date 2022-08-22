package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/kjurkovic/airtable/services/service-audit/datastore"
	"github.com/kjurkovic/airtable/services/service-audit/models"
)

func (service *AuditService) WriteLog(rw http.ResponseWriter, r *http.Request) {
	event := &models.Event{}
	err := event.Deserialize(r.Body)

	if err != nil {
		fmt.Print(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	event.Id = uuid.New()
	event.CreatedAt = time.Now()
	event.UpdatedAt = time.Now()

	err = datastore.AuditDao.AddLog(event)

	if err != nil {
		fmt.Print(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
