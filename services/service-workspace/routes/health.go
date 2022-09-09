package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kjurkovic/airtable/service/workspace/config"
	"github.com/kjurkovic/airtable/service/workspace/util"
)

type HealthCheck struct {
	Log    *util.Logger
	Config *config.Config
}

func (routes *HealthCheck) Prefix() string {
	return "/health"
}

func (routes *HealthCheck) RouteMiddleware() []mux.MiddlewareFunc {
	return []mux.MiddlewareFunc{}
}

func (routes *HealthCheck) Setup(router *mux.Router) {
	routes.Log.Info("Initializing health service routes")

	// GET router
	get := router.Methods(http.MethodGet).Subrouter()
	get.Handle("", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
}
