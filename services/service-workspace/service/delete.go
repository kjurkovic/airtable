package service

import (
	"net/http"

	"github.com/kjurkovic/airtable/service/workspace/datastore"
	"github.com/kjurkovic/airtable/service/workspace/middleware"
	"github.com/kjurkovic/airtable/service/workspace/models"
	"github.com/kjurkovic/airtable/service/workspace/wrappers"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func (service *WorkspaceService) DeleteWorkspace(rw http.ResponseWriter, r *http.Request) {
	id := uuid.MustParse(mux.Vars(r)["id"])
	service.Log.Info("DELETE Workspace#id: %s", id.String())
	rw.Header().Set("Content-Type", "application/json")

	claims := r.Context().Value(middleware.KeyClaims{}).(*models.Claims)
	affectedRows, err := datastore.WorkspaceDao.Delete(id, claims.UserId)

	if err == gorm.ErrRecordNotFound || affectedRows == 0 {
		http.Error(rw, "Workspace Not Found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(rw, "Workspace Not Found", http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusNoContent)

	wrappers.Audit.SendEvent(claims.UserId, id.String(), wrappers.WorkspaceDeleted)
}
