package service

import (
	"net/http"

	"github.com/kjurkovic/airtable/service/workspace/datastore"
	"github.com/kjurkovic/airtable/service/workspace/middleware"
	"github.com/kjurkovic/airtable/service/workspace/models"
)

func (service *WorkspaceService) GetWorkspaces(rw http.ResponseWriter, r *http.Request) {
	service.Log.Info("GET Workspaces")
	rw.Header().Set("Content-Type", "application/json")

	claims := r.Context().Value(middleware.KeyClaims{}).(*models.Claims)
	workspaces, _ := datastore.WorkspaceDao.GetAll(claims.UserId)
	err := workspaces.Serialize(rw)

	if err != nil {
		http.Error(rw, "Serialization error", http.StatusInternalServerError)
		return
	}
}
