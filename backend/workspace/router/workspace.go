package router

import (
	"fmt"
	"log"
	"net/http"
	"workspace/config"
	"workspace/database"
	"workspace/handlers"

	"github.com/gorilla/mux"
)

var workspaceIdRegex = "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}"

type WorkspaceRouter struct {
	log    *log.Logger
	config *config.Authorization
}

func (wr *WorkspaceRouter) pathPrefix() string {
	return "/workspace"
}

func (wr *WorkspaceRouter) setup(router *mux.Router) {
	wh := handlers.Workspaces(wr.log, &database.WorkspaceDao{Conn: *database.Instance}, wr.config)

	// GET router
	get := router.Methods(http.MethodGet).Subrouter()
	get.HandleFunc("", wh.GetWorkspaces)
	get.Use(wh.MiddlewareAuthorization)

	// PUT router
	put := router.Methods(http.MethodPut).Subrouter()
	put.HandleFunc(fmt.Sprintf("/{id:%s}", workspaceIdRegex), wh.UpdateWorkspace)
	put.Use(wh.MiddlewareAuthorization, wh.MiddlewareValidateWorkspace)

	// POST router
	post := router.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("", wh.AddWorkspace)
	post.Use(wh.MiddlewareAuthorization, wh.MiddlewareValidateWorkspace)

	// DELETE router
	delete := router.Methods(http.MethodDelete).Subrouter()
	delete.HandleFunc(fmt.Sprintf("/{id:%s}", workspaceIdRegex), wh.DeleteWorkspace)
	delete.Use(wh.MiddlewareAuthorization)
}
