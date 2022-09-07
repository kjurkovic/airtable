package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kjurkovic/airtable/services/service-audit/config"
	"github.com/kjurkovic/airtable/services/service-audit/middleware"
	"github.com/kjurkovic/airtable/services/service-audit/service"
	"github.com/kjurkovic/airtable/services/service-audit/util"
)

var userIdRegex = "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}"

type AuditRoutes struct {
	Log    *util.Logger
	Config *config.Config
}

func (routes *AuditRoutes) Prefix() string {
	return ""
}

func (routes *AuditRoutes) RouteMiddleware() []mux.MiddlewareFunc {
	return []mux.MiddlewareFunc{middleware.JsonResponseMiddleware}
}

func (routes *AuditRoutes) Setup(router *mux.Router) {
	routes.Log.Info("Initializing audit log service routes")
	service := &service.AuditService{
		Log:    routes.Log,
		Config: routes.Config,
	}

	post := router.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("/logs", service.WriteLog)

	get := router.Methods(http.MethodGet).Subrouter()
	get.Use(middleware.AuthMiddleware)
	get.HandleFunc(fmt.Sprintf("/audit/{id:%s}", userIdRegex), service.GetUserLogs)
}
