package service

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/kjurkovic/airtable/service/workspace/datastore"
	"github.com/kjurkovic/airtable/service/workspace/middleware"
	"github.com/kjurkovic/airtable/service/workspace/models"
	"github.com/kjurkovic/airtable/service/workspace/util"
	"github.com/kjurkovic/airtable/service/workspace/wrappers"
	"gorm.io/gorm"
)

func (service *WorkspaceService) UpdateWorkspace(rw http.ResponseWriter, r *http.Request) {
	id := uuid.MustParse(mux.Vars(r)["id"])
	service.Log.Info("PUT Workspace#id: %s", id.String())

	rw.Header().Set("Content-Type", "application/json")

	workspace := r.Context().Value(middleware.KeyWorkspace{}).(models.Workspace)
	claims := r.Context().Value(middleware.KeyClaims{}).(*models.Claims)

	response, err := datastore.WorkspaceDao.Update(id, claims.UserId, &workspace)

	if err == gorm.ErrRecordNotFound {
		http.Error(rw, "Workspace Not Found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(rw, "Workspace Not Found", http.StatusInternalServerError)
		return
	}

	err = response.Serialize(rw)

	if err != nil {
		http.Error(rw, "Serialization error", http.StatusInternalServerError)
		return
	} else {
		auditObj, _ := util.ToJson(response)
		wrappers.Audit.SendEvent(claims.UserId, auditObj, wrappers.WorkspaceModified)
	}
}
