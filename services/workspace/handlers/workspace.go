package handlers

import (
	"fmt"
	"log"
	"net/http"
	"workspace/config"
	"workspace/database"
	"workspace/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// types
type WorkspaceHandler struct {
	logger *log.Logger
	dao    *database.WorkspaceDao
	config *config.Authorization
}

type KeyWorkspace struct{}

func Workspaces(l *log.Logger, dao *database.WorkspaceDao, config *config.Authorization) *WorkspaceHandler {
	return &WorkspaceHandler{
		logger: l,
		dao:    dao,
		config: config,
	}
}

// HTTP handlers

// swagger:route POST /workspace workspace listWorkspaces
// Returns a list of workspaces
// responses:
// 	200: workspacesResponse

// GetWorkspaces returns a list of workspaces for user
func (handler *WorkspaceHandler) AddWorkspace(rw http.ResponseWriter, r *http.Request) {
	handler.logger.Print("POST Workspaces")

	rw.Header().Set("Content-Type", "application/json")

	workspace := r.Context().Value(KeyWorkspace{}).(models.Workspace)
	user := r.Context().Value(models.UserKey{}).(*models.User)
	workspace.ID = uuid.New()
	workspace.UserId = user.ID
	fmt.Printf("Received object: %v", workspace)
	_, err := handler.dao.Insert(&workspace)

	if err != nil {
		http.Error(rw, "Unable to add workspace", http.StatusInternalServerError)
		return
	}

	err = workspace.Serialize(rw)

	if err != nil {
		http.Error(rw, "Serialization error", http.StatusInternalServerError)
		return
	}
}

func (handler *WorkspaceHandler) UpdateWorkspace(rw http.ResponseWriter, r *http.Request) {
	id := uuid.MustParse(mux.Vars(r)["id"])
	handler.logger.Printf("PUT Workspace#id: %s", id.String())

	rw.Header().Set("Content-Type", "application/json")

	workspace := r.Context().Value(KeyWorkspace{}).(models.Workspace)
	user := r.Context().Value(models.UserKey{}).(*models.User)

	response, err := handler.dao.Update(id, user.ID, &workspace)

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
	}
}

// swagger:route GET /workspace workspace listWorkspaces
// Returns a list of workspaces
// responses:
// 	200: workspacesResponse

// GetWorkspaces returns a list of workspaces for user
func (handler *WorkspaceHandler) GetWorkspaces(rw http.ResponseWriter, r *http.Request) {
	handler.logger.Print("GET Workspaces")
	rw.Header().Set("Content-Type", "application/json")
	user := r.Context().Value(models.UserKey{}).(*models.User)
	workspaces, _ := handler.dao.GetAll(user.ID)
	err := workspaces.Serialize(rw)

	if err != nil {
		http.Error(rw, "Serialization error", http.StatusInternalServerError)
		return
	}
}

// swagger:route DELETE /workspace/{id} workspace deleteWorkspace
// responses:
// 	201: noContent

// Deletes Workspace from the database
func (handler *WorkspaceHandler) DeleteWorkspace(rw http.ResponseWriter, r *http.Request) {
	id := uuid.MustParse(mux.Vars(r)["id"])
	handler.logger.Printf("DELETE Workspace#id: %s", id.String())
	rw.Header().Set("Content-Type", "application/json")

	user := r.Context().Value(models.UserKey{}).(*models.User)
	affectedRows, err := handler.dao.Delete(id, user.ID)

	if err == gorm.ErrRecordNotFound || affectedRows == 0 {
		http.Error(rw, "Workspace Not Found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(rw, "Workspace Not Found", http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}
