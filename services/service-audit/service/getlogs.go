package service

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/kjurkovic/airtable/services/service-audit/datastore"
	"github.com/kjurkovic/airtable/services/service-audit/util"
)

func (service *AuditService) GetUserLogs(rw http.ResponseWriter, r *http.Request) {
	service.Log.Info("Audit service GET user logs %s", mux.Vars(r)["id"])
	id := uuid.MustParse(mux.Vars(r)["id"])

	page, pageSize := util.GetPaginationParams(r)
	auditType := r.URL.Query().Get("type")

	fmt.Println(auditType)

	content, err := datastore.AuditDao.GetUserLogs(id, auditType, page, pageSize)

	if err != nil {
		service.Log.Error(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	content.Serialize(rw)
}
