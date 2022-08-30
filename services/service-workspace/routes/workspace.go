package routes

import (
	"fmt"
	"net/http"

	"github.com/kjurkovic/airtable/service/workspace/config"
	"github.com/kjurkovic/airtable/service/workspace/middleware"
	"github.com/kjurkovic/airtable/service/workspace/service"
	"github.com/kjurkovic/airtable/service/workspace/util"

	"github.com/gorilla/mux"
)

var workspaceIdRegex = "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}"

type WorkspaceRoutes struct {
	Log    *util.Logger
	Config *config.Config
}

func (workspaceRoutes *WorkspaceRoutes) Prefix() string {
	return "/workspace"
}

func (workspaceRoutes *WorkspaceRoutes) RouteMiddleware() []mux.MiddlewareFunc {
	return []mux.MiddlewareFunc{middleware.AuthMiddleware, middleware.JsonResponseMiddleware}
}

func (workspaceRoutes *WorkspaceRoutes) Setup(router *mux.Router) {
	service := service.WorkspaceService{
		Log:    workspaceRoutes.Log,
		Config: workspaceRoutes.Config,
	}

	// GET router
	get := router.Methods(http.MethodGet).Subrouter()
	get.HandleFunc("", service.GetWorkspaces)

	// PUT router
	put := router.Methods(http.MethodPut).Subrouter()
	put.HandleFunc(fmt.Sprintf("/{id:%s}", workspaceIdRegex), service.UpdateWorkspace)
	put.Use(middleware.MiddlewareValidateWorkspace)

	// POST router
	post := router.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("", service.AddWorkspace)
	post.Use(middleware.MiddlewareValidateWorkspace)

	// DELETE router
	delete := router.Methods(http.MethodDelete).Subrouter()
	delete.HandleFunc(fmt.Sprintf("/{id:%s}", workspaceIdRegex), service.DeleteWorkspace)
}
