package service

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/kjurkovic/airtable/service/workspace/datastore"
	"github.com/kjurkovic/airtable/service/workspace/middleware"
	"github.com/kjurkovic/airtable/service/workspace/models"
	"github.com/kjurkovic/airtable/service/workspace/util"
	"github.com/kjurkovic/airtable/service/workspace/wrappers"
)

func (service *WorkspaceService) AddWorkspace(rw http.ResponseWriter, r *http.Request) {
	service.Log.Info("POST Workspaces")

	rw.Header().Set("Content-Type", "application/json")

	workspace := r.Context().Value(middleware.KeyWorkspace{}).(models.Workspace)
	claims := r.Context().Value(middleware.KeyClaims{}).(*models.Claims)
	workspace.ID = uuid.New()
	workspace.UserId = claims.UserId
	fmt.Printf("Received object: %v", workspace)
	_, err := datastore.WorkspaceDao.Insert(&workspace)

	if err != nil {
		http.Error(rw, "Unable to add workspace", http.StatusInternalServerError)
		return
	}

	err = workspace.Serialize(rw)

	if err != nil {
		http.Error(rw, "Serialization error", http.StatusInternalServerError)
		return
	} else {
		auditObj, _ := util.ToJson(workspace)
		wrappers.Audit.SendEvent(claims.UserId, auditObj, wrappers.WorkspaceCreated)
	}
}
